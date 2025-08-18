package core

import "context"

type IPermissionEntity interface {
	Id() string
	EntityType() PermissionEntityType
}

type IPermissionRequired interface {
	IsManagement() bool
	OrgId() string
	Key() string
	RequiredResourcesType() []string
	RequiredResourcesId() []string
	Action() string // used with PermissionDeniedError
}

type IAccessManagementPC interface {
	CheckPermission(context.Context, IPermissionEntity, IPermissionRequired) error
	CheckSelfPermission(currentUserId string, userId string, action string) error
}
