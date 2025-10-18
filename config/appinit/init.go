package appinit

import (
	"context"

	"github.com/Adgytec/adgytec-flow/config/app"
	"github.com/Adgytec/adgytec-flow/config/database"
	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/services/user"
	"github.com/Adgytec/adgytec-flow/utils/actor"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/rs/zerolog/log"
)

type permissionType string

const (
	permissionTypeManagement  permissionType = "management"
	permissionTypeApplication permissionType = "application"
)

type serviceFactory func() (db.AddServicesIntoStagingParams, []db.AddManagementPermissionsIntoStagingParams, []db.AddApplicationPermissionsIntoStagingParams, []db.AddServiceRestrictionIntoStagingParams)

var appServices = []serviceFactory{
	iam.InitIAMService,
	user.InitUserService,
}

func EnsureServicesInitialization(appConfig app.App) error {
	log.Info().Msg("Ensuring all application services are initialized")

	// system context
	systemCtx := actor.NewSystemActorContext(context.Background())

	var allServicesDetails []db.AddServicesIntoStagingParams
	var allManagementPermissions []db.AddManagementPermissionsIntoStagingParams
	var allApplicationPermissions []db.AddApplicationPermissionsIntoStagingParams
	var allServicesRestrictions []db.AddServiceRestrictionIntoStagingParams

	for _, factory := range appServices {
		details, managementPermissions, applicationPermissions, restrictions := factory()

		allServicesDetails = append(allServicesDetails, details)
		allManagementPermissions = append(allManagementPermissions, managementPermissions...)
		allApplicationPermissions = append(allApplicationPermissions, applicationPermissions...)
		allServicesRestrictions = append(allServicesRestrictions, restrictions...)
	}

	if err := genericAdd(
		systemCtx,
		appConfig.Database(),
		allServicesDetails,
		nil, // no ID modification needed for services
		func(ctx context.Context, q *db.Queries) error { return q.NewServiceStagingTable(ctx) },
		func(ctx context.Context, q *db.Queries, items []db.AddServicesIntoStagingParams) (int64, error) {
			return q.AddServicesIntoStaging(ctx, items)
		},
		func(ctx context.Context, q *db.Queries) error { return q.UpsertServicesFromStaging(ctx) },
	); err != nil {
		return &AddingServiceDetailsError{
			cause: err,
		}
	}

	if err := genericAdd(
		systemCtx,
		appConfig.Database(),
		allManagementPermissions,
		func(p *db.AddManagementPermissionsIntoStagingParams) {
			p.ID = core.GetIDFromPayload([]byte(p.Key))
		},
		func(ctx context.Context, q *db.Queries) error { return q.NewManagementPermissionStagingTable(ctx) },
		func(ctx context.Context, q *db.Queries, items []db.AddManagementPermissionsIntoStagingParams) (int64, error) {
			return q.AddManagementPermissionsIntoStaging(ctx, items)
		},
		func(ctx context.Context, q *db.Queries) error { return q.UpsertManagementPermissionsFromStaging(ctx) },
	); err != nil {
		return &AddingPermissionError{
			permissionType: permissionTypeManagement,
			cause:          err,
		}
	}

	if err := genericAdd(
		systemCtx,
		appConfig.Database(),
		allApplicationPermissions,
		func(p *db.AddApplicationPermissionsIntoStagingParams) {
			p.ID = core.GetIDFromPayload([]byte(p.Key))
		},
		func(ctx context.Context, q *db.Queries) error { return q.NewApplicationPermissionStagingTable(ctx) },
		func(ctx context.Context, q *db.Queries, items []db.AddApplicationPermissionsIntoStagingParams) (int64, error) {
			return q.AddApplicationPermissionsIntoStaging(ctx, items)
		},
		func(ctx context.Context, q *db.Queries) error { return q.UpsertApplicationPermissionsFromStaging(ctx) },
	); err != nil {
		return &AddingPermissionError{
			permissionType: permissionTypeApplication,
			cause:          err,
		}
	}

	if err := genericAdd(
		systemCtx,
		appConfig.Database(),
		allServicesRestrictions,
		func(p *db.AddServiceRestrictionIntoStagingParams) {
			payload := append(p.ServiceID[:], []byte(p.Name)...)
			p.ID = core.GetIDFromPayload(payload)
		},
		func(ctx context.Context, q *db.Queries) error { return q.NewServiceRestrictionsStagingTable(ctx) },
		func(ctx context.Context, q *db.Queries, items []db.AddServiceRestrictionIntoStagingParams) (int64, error) {
			return q.AddServiceRestrictionIntoStaging(ctx, items)
		},
		func(ctx context.Context, q *db.Queries) error { return q.UpsertServiceRestrictionsFromStaging(ctx) },
	); err != nil {
		return &AddServiceRestrictionsError{
			cause: err,
		}
	}

	return nil
}

// genericAdd handles services or permissions in staging tables
func genericAdd[T any](
	ctx context.Context,
	dbConn database.Database,
	items []T,
	idFunc func(*T), // optional, can be nil
	createStaging func(ctx context.Context, q *db.Queries) error,
	copyIntoStaging func(ctx context.Context, q *db.Queries, items []T) (int64, error),
	upsert func(ctx context.Context, q *db.Queries) error,
) error {

	// optionally set IDs
	if idFunc != nil {
		for i := range items {
			idFunc(&items[i])
		}
	}

	qtx, tx, txErr := dbConn.WithTransaction(ctx)
	if txErr != nil {
		return txErr
	}
	defer tx.Rollback(ctx)

	// create temp table
	if err := createStaging(ctx, qtx.Queries()); err != nil {
		return err
	}

	// copy into staging
	if _, err := copyIntoStaging(ctx, qtx.Queries(), items); err != nil {
		return err
	}

	// upsert into main table
	if err := upsert(ctx, qtx.Queries()); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
