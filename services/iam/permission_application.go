package iam

import (
	"fmt"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

var applicationPermissions = []db.AddApplicationPermissionsIntoStagingParams{
	assignApplicationPermission,
	removeApplicationPermission,
	listApplicationPermission,
}

var assignApplicationPermission = db.AddApplicationPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:assign:permission", iamServiceDetails.Name),
	ServiceID: iamServiceDetails.ID,
	Name:      "Assign Permission",
	Description: pointer.New(`
### Assign Permission

Grants the ability to assign permissions to any user or group.`),
	RequiredResources: []string{},
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}

var removeApplicationPermission = db.AddApplicationPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:remove:permission", iamServiceDetails.Name),
	ServiceID: iamServiceDetails.ID,
	Name:      "Remove Permission",
	Description: pointer.New(`
### Remove Permission

Grants the ability to remove permissions from any user or group.`),
	RequiredResources: []string{},
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}

var listApplicationPermission = db.AddApplicationPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:list:permission", iamServiceDetails.Name),
	ServiceID: iamServiceDetails.ID,
	Name:      "List Permission",
	Description: pointer.New(`
### List Permission

Grants the ability to list permissions to any user or group.`),
	RequiredResources: []string{},
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}
