package appinit

import (
	"context"

	"github.com/Adgytec/adgytec-flow/config/app"
	"github.com/Adgytec/adgytec-flow/config/database"
	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/services/user"
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

	var allServicesDetails []db.AddServicesIntoStagingParams
	var allManagementPermissions []db.AddManagementPermissionsIntoStagingParams
	var allApplicaitonPermissions []db.AddApplicationPermissionsIntoStagingParams

	for _, factory := range appServices {
		details, managementPermissions, applicationPermissions := factory()

		allServicesDetails = append(allServicesDetails, details)
		allManagementPermissions = append(allManagementPermissions, managementPermissions...)
		allApplicaitonPermissions = append(allApplicaitonPermissions, applicationPermissions...)
	}

	if err := addServiceDetails(context.Background(), appConfig.Database(), allServicesDetails); err != nil {
		return &AddingServiceDetailsError{
			cause: err,
		}
	}

	// add management permissions
	_, managementErr := appConfig.Database().Queries().AddManagementPermissionsIntoStaging(context.Background(), allManagementPermissions)
	if managementErr != nil {
		return &AddingPermissionError{
			permissionType: permissionTypeManagement,
			cause:          managementErr,
		}
	}

	// add application permissions
	_, applicationErr := appConfig.Database().Queries().AddApplicationPermissionsIntoStaging(context.Background(), allApplicaitonPermissions)
	if applicationErr != nil {
		return &AddingPermissionError{
			permissionType: permissionTypeApplication,
			cause:          applicationErr,
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
