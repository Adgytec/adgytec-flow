package core

import "github.com/google/uuid"

type PermissionEntity struct {
	ID         uuid.UUID
	EntityType ActorType
}

type PermissionRequired struct {
	Key                 string
	RequiredResources   []string
	Management          bool
	OrgId               string
	RequiredResourcesId []string
	Action              string
}
