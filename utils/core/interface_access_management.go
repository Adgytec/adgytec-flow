package core

import (
	"context"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
)

type IAccessManagementPC interface {
	// CheckPermission returns nil if any of the IPermissionRequired is successfully resolved
	CheckPermission(context.Context, []IPermissionRequired) error
}

// IPermissionRequired provides common interface for all the permission types for easy resolution
type IPermissionRequired interface {
	GetPermissionKey() string
	GetPermissionType() PermissionType
	GetPermissionActorType() db_actions.GlobalAssignableActorType
	GetPermissionRequiredResources() PermissionRequiredResources

	// GetActionName returns current service action name
	GetActionName() string
}
