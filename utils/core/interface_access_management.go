package core

import (
	"context"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
)

type AccessManagementPC interface {
	// CheckPermission checks a single permission and returns nil if it is granted.
	CheckPermission(context.Context, PermissionProvider) error

	// CheckPermissions returns nil if any of the provided permissions are granted.
	CheckPermissions(context.Context, []PermissionProvider) error
}

// PermissionProvider provides common interface for all the permission types for easy resolution
type PermissionProvider interface {
	GetPermissionKey() string
	GetPermissionType() PermissionType
	GetPermissionActorType() db_actions.GlobalAssignableActorType
	GetPermissionRequiredResources() PermissionRequiredResources
}
