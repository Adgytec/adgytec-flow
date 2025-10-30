package iam

import (
	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/google/uuid"
)

// PermissionProvider provides common interface for all the permission types for easy resolution
type PermissionProvider interface {
	GetPermissionKey() string
	GetPermissionType() PermissionType
	GetPermissionActorType() db.GlobalAssignableActorType
	GetPermissionRequiredResources() PermissionRequiredResources
	SystemAllowed() bool
	AllowSystem() PermissionProvider
}

// PermissionRequiredResources defines the required resources for PermissionEntity for successfull resolution
type PermissionRequiredResources struct {
	OrgID                 *uuid.UUID
	ProjectID             *uuid.UUID
	UserID                *uuid.UUID
	ServiceHierarchyID    *uuid.UUID
	ServiceResourceItemID *uuid.UUID
}

// permissionEntity defines the current actor details for permission resolution
type permissionEntity struct {
	id         uuid.UUID
	entityType db.GlobalActorType
}

type permissionRequired struct {
	key                 string
	permissionType      PermissionType
	permissionActorType db.GlobalAssignableActorType
	requiredResources   PermissionRequiredResources
	systemAllowed       bool
}

func (p permissionRequired) GetPermissionKey() string {
	return p.key
}

func (p permissionRequired) GetPermissionType() PermissionType {
	return p.permissionType
}

func (p permissionRequired) GetPermissionActorType() db.GlobalAssignableActorType {
	return p.permissionActorType
}

func (p permissionRequired) GetPermissionRequiredResources() PermissionRequiredResources {
	return p.requiredResources
}

func (p permissionRequired) SystemAllowed() bool {
	return p.systemAllowed
}

func (p permissionRequired) AllowSystem() PermissionProvider {
	p.systemAllowed = true
	return p
}

// helper methods to create PermissionProvider for permission resolution

func NewPermissionRequiredFromManagementPermission(
	permission db.AddManagementPermissionsIntoStagingParams,
	requiredPermissionResources PermissionRequiredResources,
) PermissionProvider {
	return permissionRequired{
		key:                 permission.Key,
		permissionType:      PermissionTypeManagement,
		permissionActorType: permission.AssignableActor,
		requiredResources:   requiredPermissionResources,
	}
}

func NewPermissionRequiredFromApplicationPermission(
	permission db.AddApplicationPermissionsIntoStagingParams,
	requiredPermissionResources PermissionRequiredResources,
) PermissionProvider {
	return permissionRequired{
		key:                 permission.Key,
		permissionType:      PermissionTypeApplication,
		permissionActorType: permission.AssignableActor,
		requiredResources:   requiredPermissionResources,
	}
}

func NewPermissionRequiredFromSelfPermission(
	permission SelfPermissions,
	requiredPermissionResources PermissionRequiredResources,
) PermissionProvider {
	return permissionRequired{
		key:                 permission.Key,
		permissionType:      PermissionTypeSelf,
		permissionActorType: db.GlobalAssignableActorTypeUser,
		requiredResources:   requiredPermissionResources,
	}
}
