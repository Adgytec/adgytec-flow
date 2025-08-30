package core

import (
	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/google/uuid"
)

// SelfPermissions defines the permission type used to define self actions
// this is not stored in db
// SelfPermissions are not assignable to any user and are implictly available for all the users for their account actions
type SelfPermissions struct {
	Key         string
	Name        string
	Description string
}

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
