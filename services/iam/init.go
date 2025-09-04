package iam

import (
	"context"
	"log"

	"github.com/Adgytec/adgytec-flow/config/database"
	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
)

type iamServiceInit struct {
	db                     database.Database
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
	log.Printf("adding %s-service details", serviceName)
	return i.db.Queries().AddService(context.TODO(), i.serviceDetails)
}

func (i *iamServiceInit) initServiceManagementPermissions() error {
	log.Printf("adding %s-service management permissions", serviceName)

	for _, perm := range i.managementPermissions {
		perm.ID = helpers.GetIDFromPayload([]byte(perm.Key))
		if err := i.db.Queries().AddManagementPermission(context.TODO(), perm); err != nil {
			return err
		}
	}
	return nil
}

func (i *iamServiceInit) initServiceApplicationPermissions() error {
	log.Printf("adding %s-service application permissions", serviceName)
	for _, perm := range i.applicationPermissions {
		perm.ID = helpers.GetIDFromPayload([]byte(perm.Key))
		if err := i.db.Queries().AddApplicationPermission(context.TODO(), perm); err != nil {
			return err
		}
	}
	return nil
}

type iamInitParams interface {
	Database() database.Database
}

func InitIAMService(params iamInitParams) core.ServiceInit {
	return &iamServiceInit{
		db:                     params.Database(),
		serviceDetails:         iamServiceDetails,
		managementPermissions:  managementPermissions,
		applicationPermissions: applicationPermissions,
	}
}
