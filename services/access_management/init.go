package access_management

import (
	"context"
	"encoding/json"
	"log"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type accessManagementInit struct {
	db                     core.IDatabase
	serviceDetails         db_actions.AddServiceParams
	managementPermissions  []db_actions.AddManagementPermissionParams
	applicationPermissions []db_actions.AddApplicationPermissionParams
}

func (i *accessManagementInit) InitService() error {
	if err := i.initServiceDetails(); err != nil {
		return err
	}

	if err := i.initServiceManagementPermissions(); err != nil {
		return err
	}

	if err := i.initServiceApplicationPermissions(); err != nil {
		return err
	}

	return nil
}

func (i *accessManagementInit) initServiceDetails() error {
	log.Println("adding access-management service details")
	return i.db.Queries().AddService(context.TODO(), i.serviceDetails)
}

func (i *accessManagementInit) initServiceManagementPermissions() error {
	log.Println("adding access-managment management permissions")
	jsonPermissions, err := json.Marshal(i.managementPermissions)
	if err != nil {
		return err
	}

	return i.db.Queries().BatchAddManagementPermission(context.TODO(), jsonPermissions)
}

func (i *accessManagementInit) initServiceApplicationPermissions() error {
	log.Println("adding access-management application permissions.")
	jsonPermissions, err := json.Marshal(i.applicationPermissions)
	if err != nil {
		return err
	}

	return i.db.Queries().BatchAddApplicationPermission(context.TODO(), jsonPermissions)
}

type iAccessManagementInitParams interface {
	Database() core.IDatabase
}

func InitAccessManagement(params iAccessManagementInitParams) core.IServiceInit {
	return &accessManagementInit{
		db:                     params.Database(),
		serviceDetails:         accessManagementDetails,
		managementPermissions:  managementPermissions,
		applicationPermissions: applicationPermissions,
	}
}
