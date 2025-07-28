package access_management

import (
	"context"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/utils/interfaces"
)

type accessMangement struct{}

type accessManagementInit struct {
	db                     interfaces.IDatabase
	serviceDetails         db_actions.AddServiceParams
	managementPermissions  []db_actions.AddManagementPermissionsParams
	applicationPermissions []db_actions.AddApplicationPermissionsParams
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
	_, err := i.db.Queries().AddManagementPermissions(context.TODO(), i.managementPermissions)
	return err
}

func (i *accessManagementInit) initServiceApplicationPermissions() error {
	_, err := i.db.Queries().AddApplicationPermissions(context.TODO(), i.applicationPermissions)
	return err
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
