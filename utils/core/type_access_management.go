package core

import "github.com/google/uuid"

type PermissionEntityType string

const (
	PermissionEntityTypeUser   PermissionEntityType = "user"
	PermissionEntityTypeAPIKey PermissionEntityType = "api-key"
)

type PermissionEntity struct {
	ID         uuid.UUID
	EntityType PermissionEntityType
}

type PermissionRequired struct {
	Key                 string
	RequiredResources   []string
	Management          bool
	OrgId               string
	RequiredResourcesId []string
	Action              string
}
