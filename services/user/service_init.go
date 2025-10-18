package user

import (
	"github.com/Adgytec/adgytec-flow/database/db"
)

func InitUserService() (db.AddServicesIntoStagingParams, []db.AddManagementPermissionsIntoStagingParams, []db.AddApplicationPermissionsIntoStagingParams, []db.AddServiceRestrictionIntoStagingParams) {
	return userServiceDetails, managementPermissions, nil, nil
}
