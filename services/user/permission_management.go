package user

import (
	"fmt"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
)

var managementPermissions = []db_actions.AddManagementPermissionParams{
	listAllUsersPermission,
}

var listAllUsersPermission = db_actions.AddManagementPermissionParams{
	Key:       fmt.Sprintf("%s:list:users", userServiceDetails.Name),
	ServiceID: userServiceDetails.ID,
	Name:      "List All Users",
	Description: helpers.ValuePtr(`
### List All Users

Grants the ability to list all the users that are part of Adgytec studio.
*Note: This allows to view all the user regardless if they are part of any organization or management.*`),
	RequiredResources: []string{},
	AssignableActor:   db_actions.GlobalAssignableActorTypeUser,
}

var disableUserPermission = db_actions.AddManagementPermissionParams{
	Key:       fmt.Sprintf("%s:disable:users", userServiceDetails.Name),
	ServiceID: userServiceDetails.ID,
	Name:      "Disable Users",
	Description: helpers.ValuePtr(`
### Disable Users

Grants the ability to disable users access to Adgytec Studio.
*Note: This disables users globally regardless of the organization they belong to.*`),
	RequiredResources: []string{},
	AssignableActor:   db_actions.GlobalAssignableActorTypeUser,
}

var enableUserPermission = db_actions.AddManagementPermissionParams{
	Key:       fmt.Sprintf("%s:enable:users", userServiceDetails.Name),
	ServiceID: userServiceDetails.ID,
	Name:      "Enable Users",
	Description: helpers.ValuePtr(`
### Enable Users

Grants the ability to enable users access to Adgytec Studio.`),
	RequiredResources: []string{},
	AssignableActor:   db_actions.GlobalAssignableActorTypeUser,
}
