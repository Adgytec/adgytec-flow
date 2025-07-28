package access_management

import (
	"context"
	"encoding/json"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/utils/interfaces"
)

type accessMangement struct{}

type accessManagementInit struct {
	db                     interfaces.IDatabase
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
	return i.db.Queries().AddService(context.TODO(), i.serviceDetails)
}

func (i *accessManagementInit) initServiceManagementPermissions() error {
	jsonPermissions, err := json.Marshal(i.managementPermissions)
	if err != nil {
		return err
	}

	return i.db.Queries().BatchAddManagementPermission(context.TODO(), jsonPermissions)
}

func (i *accessManagementInit) initServiceApplicationPermissions() error {
	jsonPermissions, err := json.Marshal(i.applicationPermissions)
	if err != nil {
		return err
	}

	return i.db.Queries().BatchAddApplicationPermission(context.TODO(), jsonPermissions)
}

type iAccessManagementInitParams interface {
	Database() interfaces.IDatabase
}

func InitAccessManagement(params iAccessManagementInitParams) interfaces.IServiceInit {
	return &accessManagementInit{
		db:                     params.Database(),
		serviceDetails:         accessManagementDetails,
		managementPermissions:  managementPermissions,
		applicationPermissions: applicationPermissions,
	}
}
