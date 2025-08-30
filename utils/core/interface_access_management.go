package core

import (
	"context"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
)

type IAccessManagementPC interface {
	CheckPermission(context.Context, PermissionRequired) error
}

// IPermissionRequired provides common interface for all the permission types for easy resolution
type IPermissionRequired interface {
	GetPermissionKey() string
	GetPermissionType() PermissionType
	GetPermissionActorType() db_actions.GlobalAssignableActorType
	GetPermissionRequiredResources() PermissionRequiredResources
}
