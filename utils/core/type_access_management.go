package core

import (
	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/google/uuid"
)

// SelfPermissions defines the permission type used to define self actions
// this is not stored in db
// SelfPermissions are not assignable to any user and are implictly available for all the users for their account actions
type SelfPermissions struct {
	Key         string
	Name        string
	Description string
}

// PermissionType defines 'Type' of permission
// this is used during permission resolution in access managment
type PermissionType string

const (
	PermissionTypeSelf        PermissionType = "self"
	PermissionTypeManagement  PermissionType = "managment"
	PermissionTypeApplication PermissionType = "application"
)

// PermissionRequiredResources defines the required resources for PermissionEntity for successfull resolution
type PermissionRequiredResources struct {
	OrgID                 *uuid.UUID
	ProjectID             *uuid.UUID
	UserID                *uuid.UUID
	ServiceHierarchyID    *uuid.UUID
	ServiceResourceItemID *uuid.UUID
}

// PermissionEntity defines the current actor details for permission resolution
type PermissionEntity struct {
	ID         uuid.UUID
	EntityType db_actions.GlobalActorType
}

// PermissionRequired defines the permission details required for successfull resolution of permission
// this is not directly used
type PermissionRequired struct {
	key                 string
	permissionType      PermissionType
	permissionActorType db_actions.GlobalAssignableActorType
	requiredResources   PermissionRequiredResources
}

func (p PermissionRequired) PermissionKey() string {
	return p.key
}

func (p PermissionRequired) PermissionType() PermissionType {
	return p.permissionType
}

func (p PermissionRequired) PermissionActorType() db_actions.GlobalAssignableActorType {
	return p.permissionActorType
}

func (p PermissionRequired) PermissionRequiredResources() PermissionRequiredResources {
	return p.requiredResources
}
