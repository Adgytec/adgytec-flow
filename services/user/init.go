package user

import (
	"context"
	"encoding/json"
	"log"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type userServiceInit struct {
	db                    core.IDatabase
	serviceDetails        db_actions.AddServiceParams
	managementPermissions []db_actions.AddManagementPermissionParams
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
	jsonPermissions, err := json.Marshal(i.managementPermissions)
	if err != nil {
		return err
	}

	return i.db.Queries().BatchAddManagementPermission(context.TODO(), jsonPermissions)
}

type iUserServiceInitParams interface {
	Database() core.IDatabase
}

func InitUserService(params iUserServiceInitParams) core.IServiceInit {
	return &userServiceInit{
		db:                    params.Database(),
		serviceDetails:        userServiceDetails,
		managementPermissions: managementPermissions,
	}
}
