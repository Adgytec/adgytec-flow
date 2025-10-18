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

type serviceFactory func() (db.AddServicesIntoStagingParams, []db.AddManagementPermissionsIntoStagingParams, []db.AddApplicationPermissionsIntoStagingParams)

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

	for _, factory := range appServices {
		details, managementPermissions, applicationPermissions := factory()

		allServicesDetails = append(allServicesDetails, details)
		allManagementPermissions = append(allManagementPermissions, managementPermissions...)
		allApplicationPermissions = append(allApplicationPermissions, applicationPermissions...)
	}

	if err := addServiceDetails(systemCtx, appConfig.Database(), allServicesDetails); err != nil {
		return &AddingServiceDetailsError{
			cause: err,
		}
	}

	if err := addManagementPermissions(systemCtx, appConfig.Database(), allManagementPermissions); err != nil {
		return &AddingPermissionError{
			permissionType: permissionTypeManagement,
			cause:          err,
		}
	}

	if err := addApplicationPermissions(systemCtx, appConfig.Database(), allApplicationPermissions); err != nil {
		return &AddingPermissionError{
			permissionType: permissionTypeApplication,
			cause:          err,
		}
	}

	return nil
}

func addServiceDetails(ctx context.Context, dbConn database.Database, serviceDetails []db.AddServicesIntoStagingParams) error {
	qtx, tx, txErr := dbConn.WithTransaction(ctx)
	if txErr != nil {
		return txErr
	}
	defer tx.Rollback(ctx)

	// create temp table
	stagingErr := qtx.Queries().NewServiceStagingTable(ctx)
	if stagingErr != nil {
		return stagingErr
	}

	// add service details to staging
	_, serviceErr := qtx.Queries().AddServicesIntoStaging(ctx, serviceDetails)
	if serviceErr != nil {
		return serviceErr
	}

	// add to main services
	upsertErr := qtx.Queries().UpsertServicesFromStaging(ctx)
	if upsertErr != nil {
		return upsertErr
	}

	return tx.Commit(ctx)
}

func addManagementPermissions(ctx context.Context, dbConn database.Database, permissions []db.AddManagementPermissionsIntoStagingParams) error {
	for i := range permissions {
		permissions[i].ID = core.GetIDFromPayload([]byte(permissions[i].Key))
	}

	qtx, tx, txErr := dbConn.WithTransaction(ctx)
	if txErr != nil {
		return txErr
	}
	defer tx.Rollback(ctx)

	// create temp table
	stagingErr := qtx.Queries().NewManagementPermissionStagingTable(ctx)
	if stagingErr != nil {
		return stagingErr
	}

	// add service details to staging
	_, permissionErr := qtx.Queries().AddManagementPermissionsIntoStaging(ctx, permissions)
	if permissionErr != nil {
		return permissionErr
	}

	// add to main services
	upsertErr := qtx.Queries().UpsertManagementPermissionsFromStaging(ctx)
	if upsertErr != nil {
		return upsertErr
	}

	return tx.Commit(ctx)
}

func addApplicationPermissions(ctx context.Context, dbConn database.Database, permissions []db.AddApplicationPermissionsIntoStagingParams) error {
	for i := range permissions {
		permissions[i].ID = core.GetIDFromPayload([]byte(permissions[i].Key))
	}

	qtx, tx, txErr := dbConn.WithTransaction(ctx)
	if txErr != nil {
		return txErr
	}
	defer tx.Rollback(ctx)

	// create temp table
	stagingErr := qtx.Queries().NewApplicationPermissionStagingTable(ctx)
	if stagingErr != nil {
		return stagingErr
	}

	// add service details to staging
	_, permissionErr := qtx.Queries().AddApplicationPermissionsIntoStaging(ctx, permissions)
	if permissionErr != nil {
		return permissionErr
	}

	// add to main services
	upsertErr := qtx.Queries().UpsertApplicationPermissionsFromStaging(ctx)
	if upsertErr != nil {
		return upsertErr
	}

	return tx.Commit(ctx)
}
