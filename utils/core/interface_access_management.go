package core

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
	CheckPermission(IPermissionEntity, IPermissionRequired) error
	CheckSelfPermission(string, string) error
}
