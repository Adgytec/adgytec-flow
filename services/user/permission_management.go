package user

import (
	"fmt"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/jackc/pgx/v5/pgtype"
)

var managementPermissions = []db_actions.AddManagementPermissionParams{
	listAllUsersPermission,
}

var listAllUsersPermission = db_actions.AddManagementPermissionParams{
	Key:         fmt.Sprintf("%s:list:users", userServiceDetails.Name),
	ServiceName: userServiceDetails.Name,
	Name:        "List All Users",
	Description: pgtype.Text{
		String: `
### List All Users

Grants the ability to list all the users that are part of Adgytec studio.
*Note: This allows to view all the user regardless if they are part of any organization or management.*
		`,
		Valid: true,
	},
	RequiredResources: []db_actions.ManagementPermissionResourceType{},
}

var disableUserPermission = db_actions.AddManagementPermissionParams{
	Key:         fmt.Sprintf("%s:disable:users", userServiceDetails.Name),
	ServiceName: userServiceDetails.Name,
	Name:        "Disable Users",
	Description: pgtype.Text{
		String: `
### Disable Users

Grants the ability to disable users access to Adgytec Studio.
*Note: This disables users globally regardless of the organization they belong to.*
		`,
		Valid: true,
	},
	RequiredResources: []db_actions.ManagementPermissionResourceType{},
}

var enableUserPermission = db_actions.AddManagementPermissionParams{
	Key:         fmt.Sprintf("%s:enable:users", userServiceDetails.Name),
	ServiceName: userServiceDetails.Name,
	Name:        "Enable Users",
	Description: pgtype.Text{
		String: `
### Enable Users

Grants the ability to enable users access to Adgytec Studio.
		`,
		Valid: true,
	},
	RequiredResources: []db_actions.ManagementPermissionResourceType{},
}
