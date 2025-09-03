package user

import (
	"context"
	"log"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
)

type userServiceInit struct {
	db                    core.Database
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
	log.Println("adding user service details")
	return i.db.Queries().AddService(context.TODO(), i.serviceDetails)
}

func (i *userServiceInit) initServiceManagementPermissions() error {
	log.Println("adding user service management permissions")
	for _, perm := range i.managementPermissions {
		perm.ID = helpers.GetIDFromPayload([]byte(perm.Key))
		if err := i.db.Queries().AddManagementPermission(context.TODO(), perm); err != nil {
			return err
		}
	}
	return nil
}

type userServiceInitParams interface {
	Database() core.Database
}

func InitUserService(params userServiceInitParams) core.ServiceInit {
	return &userServiceInit{
		db:                    params.Database(),
		serviceDetails:        userServiceDetails,
		managementPermissions: managementPermissions,
	}
}
