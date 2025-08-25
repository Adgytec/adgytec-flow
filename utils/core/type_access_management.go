package core

import (
	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/google/uuid"
)

type PermissionEntity struct {
	ID         uuid.UUID
	EntityType db_actions.GlobalActorType
}

type PermissionRequired struct {
	Key                 string
	RequiredResources   []string
	Management          bool
	OrgId               string
	RequiredResourcesId []string
	Action              string
}
