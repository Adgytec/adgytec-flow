package user

import (
	"fmt"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

var managementPermissions = []db.AddManagementPermissionParams{
	listAllUsersPermission,
	disableUserPermission,
	enableUserPermission,
	getUserProfilePermission,
	updateUserProfilePermission,
}

var listAllUsersPermission = db.AddManagementPermissionParams{
	Key:       fmt.Sprintf("%s:list:users", userServiceDetails.Name),
	ServiceID: userServiceDetails.ID,
	Name:      "List All Users",
	Description: pointer.New(`
### List All Users

Grants the ability to list all the users that are part of Adgytec studio.
*Note: This allows to view all the user regardless if they are part of any organization or management.*`),
	RequiredResources: []string{},
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}

var disableUserPermission = db.AddManagementPermissionParams{
	Key:       fmt.Sprintf("%s:disable:users", userServiceDetails.Name),
	ServiceID: userServiceDetails.ID,
	Name:      "Disable Users",
	Description: pointer.New(`
### Disable Users

Grants the ability to disable users access to Adgytec Studio.
*Note: This disables users globally regardless of the organization they belong to.*`),
	RequiredResources: []string{},
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}

var enableUserPermission = db.AddManagementPermissionParams{
	Key:       fmt.Sprintf("%s:enable:users", userServiceDetails.Name),
	ServiceID: userServiceDetails.ID,
	Name:      "Enable Users",
	Description: pointer.New(`
### Enable Users

Grants the ability to enable users access to Adgytec Studio.`),
	RequiredResources: []string{},
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}

var getUserProfilePermission = db.AddManagementPermissionParams{
	Key:       fmt.Sprintf("%s:get:user-profile", userServiceDetails.Name),
	ServiceID: userServiceDetails.ID,
	Name:      "Get User Profile",
	Description: pointer.New(`
### Get User Profile

Grants the ability to get individual user profile details.`),
	RequiredResources: []string{},
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}

var updateUserProfilePermission = db.AddManagementPermissionParams{
	Key:       fmt.Sprintf("%s:update:user-profile", userServiceDetails.Name),
	ServiceID: userServiceDetails.ID,
	Name:      "Update User Profile",
	Description: pointer.New(`
### Update User Profile

Grants the ability to update individual user profile.
*Note: This also require 'Get User Profile' permission.*`),
	RequiredResources: []string{},
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}
