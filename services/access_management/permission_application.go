package access_management

import (
	"fmt"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
)

var applicationPermissions = []db_actions.AddApplicationPermissionParams{
	assignApplicationPermission,
	removeApplicationPermission,
	listApplicationPermission,
}

var assignApplicationPermission = db_actions.AddApplicationPermissionParams{
	Key:       fmt.Sprintf("%s:assign:permission", accessManagementDetails.Name),
	ServiceID: accessManagementDetails.ID,
	Name:      "Assign Permission",
	Description: helpers.ValuePtr(`
### Assign Permission

Grants the ability to assign permissions to any user or group.`),
	RequiredResources: []db_actions.ApplicationPermissionResourceType{},
}

var removeApplicationPermission = db_actions.AddApplicationPermissionParams{
	Key:       fmt.Sprintf("%s:remove:permission", accessManagementDetails.Name),
	ServiceID: accessManagementDetails.ID,
	Name:      "Remove Permission",
	Description: helpers.ValuePtr(`
### Remove Permission

Grants the ability to remove permissions from any user or group.`),
	RequiredResources: []db_actions.ApplicationPermissionResourceType{},
}

var listApplicationPermission = db_actions.AddApplicationPermissionParams{
	Key:       fmt.Sprintf("%s:list:permission", accessManagementDetails.Name),
	ServiceID: accessManagementDetails.ID,
	Name:      "List Permission",
	Description: helpers.ValuePtr(`
### List Permission

Grants the ability to list permissions to any user or group.`),
	RequiredResources: []db_actions.ApplicationPermissionResourceType{},
}
