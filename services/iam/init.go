package iam

import (
	"context"
	"log"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
)

type iamServiceInit struct {
	db                     core.Database
	serviceDetails         db.AddServiceParams
	managementPermissions  []db.AddManagementPermissionParams
	applicationPermissions []db.AddApplicationPermissionParams
}

func (i *iamServiceInit) InitService() error {
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

func (i *iamServiceInit) initServiceDetails() error {
	log.Println("adding access-management service details")
	return i.db.Queries().AddService(context.TODO(), i.serviceDetails)
}

func (i *iamServiceInit) initServiceManagementPermissions() error {
	log.Println("adding access-managment management permissions")

	for _, perm := range i.managementPermissions {
		perm.ID = helpers.GetIDFromPayload([]byte(perm.Key))
		if err := i.db.Queries().AddManagementPermission(context.TODO(), perm); err != nil {
			return err
		}
	}
	return nil
}

func (i *iamServiceInit) initServiceApplicationPermissions() error {
	log.Println("adding access-management application permissions.")
	for _, perm := range i.applicationPermissions {
		perm.ID = helpers.GetIDFromPayload([]byte(perm.Key))
		if err := i.db.Queries().AddApplicationPermission(context.TODO(), perm); err != nil {
			return err
		}
	}
	return nil
}

type iamInitParams interface {
	Database() core.Database
}

func InitAccessManagement(params iamInitParams) core.ServiceInit {
	return &iamServiceInit{
		db:                     params.Database(),
		serviceDetails:         iamServiceDetails,
		managementPermissions:  managementPermissions,
		applicationPermissions: applicationPermissions,
	}
}
