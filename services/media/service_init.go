package media

import (
	"context"
	"log"

	"github.com/Adgytec/adgytec-flow/config/database"
	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/services"
)

type mediaServiceInit struct {
	db                    database.Database
	serviceDetails        db.AddServiceParams
	managementPermissions []db.AddManagementPermissionParams
}

func (i *mediaServiceInit) InitService() error {
	if err := i.initServiceDetails(); err != nil {
		return err
	}

	if err := i.initServiceManagementPermissions(); err != nil {
		return err
	}

	return nil
}

func (i *mediaServiceInit) initServiceDetails() error {
	log.Printf("adding %s-service details", serviceName)
	return i.db.Queries().AddService(context.TODO(), i.serviceDetails)
}

func (i *mediaServiceInit) initServiceManagementPermissions() error {
	log.Printf("adding %s-service management permissions", serviceName)
	for _, perm := range i.managementPermissions {
		perm.ID = core.GetIDFromPayload([]byte(perm.Key))
		if err := i.db.Queries().AddManagementPermission(context.TODO(), perm); err != nil {
			return err
		}
	}
	return nil
}

type mediaServiceInitParams interface {
	Database() database.Database
}

func InitMediaService(params mediaServiceInitParams) services.Init {
	return &mediaServiceInit{
		db:                    params.Database(),
		serviceDetails:        mediaServiceDetails,
		managementPermissions: managementPermissions,
	}
}
