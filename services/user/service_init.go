package user

import (
	"github.com/Adgytec/adgytec-flow/database/db"
)

func InitUserService() (db.AddServiceDetailsParams, []db.AddManagementPermissionParams, []db.AddApplicationPermissionParams) {
	return userServiceDetails, managementPermissions, nil
}
