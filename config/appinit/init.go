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

type serviceFactory func() (db.AddServiceDetailsParams, []db.AddManagementPermissionParams, []db.AddApplicationPermissionParams)

var appServices = []serviceFactory{
	iam.InitIAMService,
	user.InitUserService,
}

func EnsureServicesInitialization(appConfig app.App) {
	log.Println("Ensuring application initialization.")
	for _, factory := range appServices {
		details, managementPermissions, applicationPermissions := factory()

		if err := appConfig.Database().Queries().AddServiceDetails(context.Background(), details); err != nil {
			log.Fatalf("failed to add service details for service %s: %v", details.Name, err)
		}

		for _, perm := range managementPermissions {
			perm.ID = core.GetIDFromPayload([]byte(perm.Key))
			if err := appConfig.Database().Queries().AddManagementPermission(context.Background(), perm); err != nil {
				log.Fatalf("failed to add management permission %s for service %s: %v", perm.Key, details.Name, err)
			}
		}

		for _, perm := range applicationPermissions {
			perm.ID = core.GetIDFromPayload([]byte(perm.Key))
			if err := appConfig.Database().Queries().AddApplicationPermission(context.Background(), perm); err != nil {
				log.Fatalf("failed to add application permission %s for service %s: %v", perm.Key, details.Name, err)
			}
		}

	}

}
