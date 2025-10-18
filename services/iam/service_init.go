package iam

import (
	"github.com/Adgytec/adgytec-flow/database/db"
)

func InitIAMService() (db.AddServicesIntoStagingParams, []db.AddManagementPermissionsIntoStagingParams, []db.AddApplicationPermissionsIntoStagingParams, []db.AddServiceRestrictionIntoStagingParams) {
	return iamServiceDetails, managementPermissions, applicationPermissions, nil
}
