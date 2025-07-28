package access_management

import (
	"fmt"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/jackc/pgx/v5/pgtype"
)

var managementPermissions = []db_actions.AddManagementPermissionParams{
	assignManagementPermission,
	removeManagementPermission,
	listManagementPermission,
}

var assignManagementPermission = db_actions.AddManagementPermissionParams{
	Key:         fmt.Sprintf("%s:assign:management-permission", accessManagementDetails.Name),
	ServiceName: accessManagementDetails.Name,
	Name:        "Assign Permission",
	Description: pgtype.Text{
		String: `
### Assign Permission

Grants the ability to assign permissions to any user or group.
		`,
		Valid: true,
	},
	RequiredResources: []db_actions.GlobalPermissionResourceType{},
}

var removeManagementPermission = db_actions.AddManagementPermissionParams{
	Key:         fmt.Sprintf("%s:remove:management-permission", accessManagementDetails.Name),
	ServiceName: accessManagementDetails.Name,
	Name:        "Remove Permission",
	Description: pgtype.Text{
		String: `
### Remove Permission

Grants the ability to remove permissions from any user or group.
		`,
		Valid: true,
	},
	RequiredResources: []db_actions.GlobalPermissionResourceType{},
}

var listManagementPermission = db_actions.AddManagementPermissionParams{
	Key:         fmt.Sprintf("%s:list:management-permission", accessManagementDetails.Name),
	ServiceName: accessManagementDetails.Name,
	Name:        "List Permission",
	Description: pgtype.Text{
		String: `
### List Permission

Grants the ability to list permissions to any user or group.
		`,
		Valid: true,
	},
	RequiredResources: []db_actions.GlobalPermissionResourceType{},
}
