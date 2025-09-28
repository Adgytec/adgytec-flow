package appinit

import (
	"context"
	"log"

	"github.com/Adgytec/adgytec-flow/config/app"
	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/services/user"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type permissionType string

const (
	permissionTypeManagement  permissionType = "management"
	permissionTypeApplication permissionType = "application"
)

type serviceFactory func() (db.AddServiceDetailsParams, []db.AddManagementPermissionParams, []db.AddApplicationPermissionParams)

var appServices = []serviceFactory{
	iam.InitIAMService,
	user.InitUserService,
}

func EnsureServicesInitialization(appConfig app.App) error {
	log.Println("Ensuring application initialization.")
	for _, factory := range appServices {
		details, managementPermissions, applicationPermissions := factory()

		if err := appConfig.Database().Queries().AddServiceDetails(context.Background(), details); err != nil {
			return &AddingServiceDetailsError{
				serviceName: details.Name,
				cause:       err,
			}
		}

		for _, perm := range managementPermissions {
			perm.ID = core.GetIDFromPayload([]byte(perm.Key))
			if err := appConfig.Database().Queries().AddManagementPermission(context.Background(), perm); err != nil {
				return &AddingPermissionError{
					serviceName:     details.Name,
					cause:           err,
					permissionKey:   perm.Key,
					permmissionType: permissionTypeManagement,
				}
			}
		}

		for _, perm := range applicationPermissions {
			perm.ID = core.GetIDFromPayload([]byte(perm.Key))
			if err := appConfig.Database().Queries().AddApplicationPermission(context.Background(), perm); err != nil {
				return &AddingPermissionError{
					serviceName:     details.Name,
					cause:           err,
					permissionKey:   perm.Key,
					permmissionType: permissionTypeApplication,
				}
			}
		}
	}

	return nil
}
