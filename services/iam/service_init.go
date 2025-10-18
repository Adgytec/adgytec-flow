package iam

import (
	"github.com/Adgytec/adgytec-flow/database/db"
)

func InitIAMService() (db.AddServicesIntoStagingParams, []db.AddManagementPermissionsIntoStagingParams, []db.AddApplicationPermissionsIntoStagingParams) {
	return iamServiceDetails, managementPermissions, applicationPermissions
}
