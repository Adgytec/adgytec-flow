package core

import (
	"context"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
)

type IAccessManagementPC interface {
	// CheckPermission checks a single permission and returns nil if it is granted.
	CheckPermission(context.Context, IPermissionRequired) error

	// CheckPermissions returns nil if any of the provided permissions are granted.
	CheckPermissions(context.Context, []IPermissionRequired) error
}

// IPermissionRequired provides common interface for all the permission types for easy resolution
type IPermissionRequired interface {
	GetPermissionKey() string
	GetPermissionType() PermissionType
	GetPermissionActorType() db_actions.GlobalAssignableActorType
	GetPermissionRequiredResources() PermissionRequiredResources
}
