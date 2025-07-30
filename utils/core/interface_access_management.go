package core

type IPermissionEntity interface {
	Id() string
	EntityType() PermissionEntityType
}

type IPermissionRequired interface {
	IsManagement() bool
	OrgId() string
	Key() string
	RequiredResourcesType() []PermissionResourceType
	RequiredResourcesId() []string
}

type IAccessManagementPC interface {
	CheckPermission(IPermissionEntity, IPermissionRequired) error
}
