package iam

import (
	"github.com/Adgytec/adgytec-flow/database/db"
)

func InitIAMService() (db.AddServiceDetailsParams, []db.AddManagementPermissionParams, []db.AddApplicationPermissionParams) {
	return iamServiceDetails, managementPermissions, applicationPermissions
}
