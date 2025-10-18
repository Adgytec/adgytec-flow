package user

import (
	"github.com/Adgytec/adgytec-flow/database/db"
)

func InitUserService() (db.AddServicesIntoStagingParams, []db.AddManagementPermissionsIntoStagingParams, []db.AddApplicationPermissionsIntoStagingParams) {
	return userServiceDetails, managementPermissions, nil
}
