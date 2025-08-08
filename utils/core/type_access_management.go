package core

import db_actions "github.com/Adgytec/adgytec-flow/database/actions"

type PermissionEntityType string

const (
	PermissionEntityTypeUser  PermissionEntityType = "user"
	PermissionEntityTypeToken PermissionEntityType = "token"
)

type permissionEntity struct {
	id         string
	entityType PermissionEntityType
}

func (p *permissionEntity) Id() string {
	return p.id
}

func (p *permissionEntity) EntityType() PermissionEntityType {
	return p.entityType
}

func CreatePermissionEntity(userId string, entityType PermissionEntityType) IPermissionEntity {
	return &permissionEntity{
		id:         userId,
		entityType: entityType,
	}
}

type permissionRequired struct {
	key                 string
	requiredResources   []string
	management          bool
	orgId               string
	requiredResourcesId []string
	action              string
}

func (p *permissionRequired) IsManagement() bool {
	return p.management
}

func (p *permissionRequired) OrgId() string {
	return p.orgId
}

func (p *permissionRequired) Key() string {
	return p.key
}

func (p *permissionRequired) RequiredResourcesType() []string {
	return p.requiredResources
}

func (p *permissionRequired) RequiredResourcesId() []string {
	return p.requiredResourcesId
}

func (p *permissionRequired) Action() string {
	return p.action
}

func CreatePermssionRequiredFromManagementPermission(permission db_actions.AddManagementPermissionParams, orgId string, requiredResourcesId []string) IPermissionRequired {
	var requiredResources []string
	for _, resourceType := range permission.RequiredResources {
		requiredResources = append(requiredResources, string(resourceType))
	}

	return &permissionRequired{
		key:                 permission.Key,
		management:          true,
		orgId:               orgId,
		requiredResources:   requiredResources,
		requiredResourcesId: requiredResourcesId,
	}
}

func CreatePermssionRequiredFromApplicationPermission(permission db_actions.AddApplicationPermissionParams, orgId string, requiredResourcesId []string) IPermissionRequired {
	var requiredResources []string
	for _, resourceType := range permission.RequiredResources {
		requiredResources = append(requiredResources, string(resourceType))
	}

	return &permissionRequired{
		key:                 permission.Key,
		management:          true,
		orgId:               orgId,
		requiredResources:   requiredResources,
		requiredResourcesId: requiredResourcesId,
	}
}
