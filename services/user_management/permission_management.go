package usermanagement

import "github.com/Adgytec/adgytec-flow/database/db"

var managementPermissions = []db.AddManagementPermissionsIntoStagingParams{
	newManagementUserPermission,
	newUserGroupPermission,
	updateUserGroupPermission,
	deleteUserGroupPermission,
	addUserInUserGroupPermission,
	removeUserFromUserGroupPermission,
}

var newManagementUserPermission = db.AddManagementPermissionsIntoStagingParams{}

var newUserGroupPermission = db.AddManagementPermissionsIntoStagingParams{}

var updateUserGroupPermission = db.AddManagementPermissionsIntoStagingParams{}

var deleteUserGroupPermission = db.AddManagementPermissionsIntoStagingParams{}

var addUserInUserGroupPermission = db.AddManagementPermissionsIntoStagingParams{}

var removeUserFromUserGroupPermission = db.AddManagementPermissionsIntoStagingParams{}
