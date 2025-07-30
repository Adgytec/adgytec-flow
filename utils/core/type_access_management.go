package core

import db_actions "github.com/Adgytec/adgytec-flow/database/actions"

type PermissionEntityType string

const (
	PermissionEntityTypeUser  PermissionEntityType = "user"
	PermissionEntityTypeToken PermissionEntityType = "token"
)

type PermissionEntity struct {
	id         string
	entityType PermissionEntityType
}

func (p *PermissionEntity) Id() string {
	return p.id
}

func (p *PermissionEntity) EntityType() PermissionEntityType {
	return p.entityType
}

type PermissionResourceType struct {
	db_actions.ManagementPermissionResourceType
	db_actions.ApplicationPermissionResourceType
}

type PermissionRequired struct {
	key                 string
	requiredResources   []PermissionResourceType
	management          bool
	orgId               string
	requiredResourcesId []string
	action              string
}

func (p *PermissionRequired) IsManagement() bool {
	return p.management
}

func (p *PermissionRequired) OrgId() string {
	return p.orgId
}

func (p *PermissionRequired) Key() string {
	return p.key
}

func (p *PermissionRequired) RequiredResourcesType() []PermissionResourceType {
	return p.requiredResources
}

func (p *PermissionRequired) RequiredResourcesId() []string {
	return p.requiredResourcesId
}

func (p *PermissionRequired) Action() string {
	return p.action
}
