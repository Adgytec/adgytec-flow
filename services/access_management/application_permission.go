package access_management

import (
	"fmt"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/utils/types"
	"github.com/jackc/pgx/v5/pgtype"
)

var applicationPermissions = []types.Permission{
	assignApplicationPermission,
	removeApplicationPermission,
	listApplicationPermission,
}

var assignApplicationPermission = types.Permission{
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
	RequiredResources: []db_actions.GlobalPermissionResourceType{},
}

var removeApplicationPermission = types.Permission{
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
	RequiredResources: []db_actions.GlobalPermissionResourceType{},
}

var listApplicationPermission = types.Permission{
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
	RequiredResources: []db_actions.GlobalPermissionResourceType{},
}
