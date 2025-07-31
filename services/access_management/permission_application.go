package access_management

import (
	"fmt"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/jackc/pgx/v5/pgtype"
)

var applicationPermissions = []db_actions.AddApplicationPermissionParams{
	assignApplicationPermission,
	removeApplicationPermission,
	listApplicationPermission,
}

var assignApplicationPermission = db_actions.AddApplicationPermissionParams{
	Key:         fmt.Sprintf("%s:assign:permission", accessManagementDetails.Name),
	ServiceName: accessManagementDetails.Name,
	Name:        "Assign Permission",
	Description: pgtype.Text{
		String: `
### Assign Permission

Grants the ability to assign permissions to any user or group.
		`,
		Valid: true,
	},
	RequiredResources: []db_actions.ApplicationPermissionResourceType{},
}

var removeApplicationPermission = db_actions.AddApplicationPermissionParams{
	Key:         fmt.Sprintf("%s:remove:permission", accessManagementDetails.Name),
	ServiceName: accessManagementDetails.Name,
	Name:        "Remove Permission",
	Description: pgtype.Text{
		String: `
### Remove Permission

Grants the ability to remove permissions from any user or group.
		`,
		Valid: true,
	},
	RequiredResources: []db_actions.ApplicationPermissionResourceType{},
}

var listApplicationPermission = db_actions.AddApplicationPermissionParams{
	Key:         fmt.Sprintf("%s:list:permission", accessManagementDetails.Name),
	ServiceName: accessManagementDetails.Name,
	Name:        "List Permission",
	Description: pgtype.Text{
		String: `
### List Permission

Grants the ability to list permissions to any user or group.
		`,
		Valid: true,
	},
	RequiredResources: []db_actions.ApplicationPermissionResourceType{},
}
