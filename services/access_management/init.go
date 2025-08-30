package access_management

import (
	"context"
	"log"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
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

	for _, perm := range i.managementPermissions {
		perm.ID = helpers.GetIDFromPayload([]byte(perm.Name))
		if err := i.db.Queries().AddManagementPermission(context.TODO(), perm); err != nil {
			return err
		}
	}
	return nil
}

func (i *accessManagementInit) initServiceApplicationPermissions() error {
	log.Println("adding access-management application permissions.")
	for _, perm := range i.applicationPermissions {
		perm.ID = helpers.GetIDFromPayload([]byte(perm.Name))
		if err := i.db.Queries().AddApplicationPermission(context.TODO(), perm); err != nil {
			return err
		}
	}
	return nil
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
