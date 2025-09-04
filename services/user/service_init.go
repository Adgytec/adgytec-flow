package user

import (
	"context"
	"log"

	"github.com/Adgytec/adgytec-flow/config/database"
	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
)

type userServiceInit struct {
	db                    database.Database
	serviceDetails        db.AddServiceParams
	managementPermissions []db.AddManagementPermissionParams
}

func (i *userServiceInit) InitService() error {
	if err := i.initServiceDetails(); err != nil {
		return err
	}

	if err := i.initServiceManagementPermissions(); err != nil {
		return err
	}

	return nil
}

func (i *userServiceInit) initServiceDetails() error {
	log.Printf("adding %s-service details", serviceName)
	return i.db.Queries().AddService(context.TODO(), i.serviceDetails)
}

func (i *userServiceInit) initServiceManagementPermissions() error {
	log.Printf("adding %s-service management permissions", serviceName)
	for _, perm := range i.managementPermissions {
		perm.ID = helpers.GetIDFromPayload([]byte(perm.Key))
		if err := i.db.Queries().AddManagementPermission(context.TODO(), perm); err != nil {
			return err
		}
	}
	return nil
}

type userServiceInitParams interface {
	Database() database.Database
}

func InitUserService(params userServiceInitParams) core.ServiceInit {
	return &userServiceInit{
		db:                    params.Database(),
		serviceDetails:        userServiceDetails,
		managementPermissions: managementPermissions,
	}
}
