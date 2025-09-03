package iam

import (
	"fmt"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

var managementPermissions = []db.AddManagementPermissionParams{
	assignManagementPermission,
	removeManagementPermission,
	listManagementPermission,
}

var assignManagementPermission = db.AddManagementPermissionParams{
	Key:       fmt.Sprintf("%s:assign:management-permission", accessManagementDetails.Name),
	ServiceID: accessManagementDetails.ID,
	Name:      "Assign Permission",
	Description: pointer.New(`
### Assign Permission

Grants the ability to assign permissions to any user or group.`),
	RequiredResources: []string{
		string(db.ManagementPermissionResourceTypeOrganization),
	},
	AssignableActor: db.GlobalAssignableActorTypeUser,
}

var removeManagementPermission = db.AddManagementPermissionParams{
	Key:       fmt.Sprintf("%s:remove:management-permission", accessManagementDetails.Name),
	ServiceID: accessManagementDetails.ID,
	Name:      "Remove Permission",
	Description: pointer.New(`
### Remove Permission

Grants the ability to remove permissions from any user or group.`),
	RequiredResources: []string{
		string(db.ManagementPermissionResourceTypeOrganization),
	},
	AssignableActor: db.GlobalAssignableActorTypeUser,
}

var listManagementPermission = db.AddManagementPermissionParams{
	Key:       fmt.Sprintf("%s:list:management-permission", accessManagementDetails.Name),
	ServiceID: accessManagementDetails.ID,
	Name:      "List Permission",
	Description: pointer.New(`
### List Permission

Grants the ability to list permissions to any user or group.`),

	RequiredResources: []string{
		string(db.ManagementPermissionResourceTypeOrganization),
	},
	AssignableActor: db.GlobalAssignableActorTypeUser,
}
