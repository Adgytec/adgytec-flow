package core

import (
	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/google/uuid"
)

// SelfPermissions defines the permission type used to define self actions
// this is not stored in db
// SelfPermissions are not assignable to any user and are implictly available for all the users for their account actions
type SelfPermissions struct {
	Key  string
	Name string
}

// PermissionType defines 'Type' of permission
// this is used during permission resolution in access managment
type PermissionType string

const (
	PermissionTypeSelf        PermissionType = "self"
	PermissionTypeManagement  PermissionType = "management"
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
	Key                 string
	PermissionType      PermissionType
	PermissionActorType db_actions.GlobalAssignableActorType
	RequiredResources   PermissionRequiredResources
}

func (p PermissionRequired) GetPermissionKey() string {
	return p.Key
}

func (p PermissionRequired) GetPermissionType() PermissionType {
	return p.PermissionType
}

func (p PermissionRequired) GetPermissionActorType() db_actions.GlobalAssignableActorType {
	return p.PermissionActorType
}

func (p PermissionRequired) GetPermissionRequiredResources() PermissionRequiredResources {
	return p.RequiredResources
}
