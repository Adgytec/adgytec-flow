package helpers

import (
	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

// helper methods to create core.IPermissionRequired for permission resolution

func NewPermissionRequiredFromManagementPermission(
	permission db_actions.AddManagementPermissionParams,
	requiredPermissionResources core.PermissionRequiredResources,
) core.PermissionProvider {
	return core.PermissionRequired{
		Key:                 permission.Key,
		PermissionType:      core.PermissionTypeManagement,
		PermissionActorType: permission.AssignableActor,
		RequiredResources:   requiredPermissionResources,
	}
}

func NewPermissionRequiredFromApplicationPermission(
	permission db_actions.AddApplicationPermissionParams,
	requiredPermissionResources core.PermissionRequiredResources,
) core.PermissionProvider {
	return core.PermissionRequired{
		Key:                 permission.Key,
		PermissionType:      core.PermissionTypeApplication,
		PermissionActorType: permission.AssignableActor,
		RequiredResources:   requiredPermissionResources,
	}
}

func NewPermissionRequiredFromSelfPermission(
	permission core.SelfPermissions,
	requiredPermissionResources core.PermissionRequiredResources,
) core.PermissionProvider {
	return core.PermissionRequired{
		Key:                 permission.Key,
		PermissionType:      core.PermissionTypeSelf,
		PermissionActorType: db_actions.GlobalAssignableActorTypeUser,
		RequiredResources:   requiredPermissionResources,
	}
}
