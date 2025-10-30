package actor

import (
	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/google/uuid"
)

type ActorKey string

const (
	ActorKeyType ActorKey = "actor-type"
	ActorKeyID   ActorKey = "actor-id"
)

var (
	systemActorID = uuid.Nil
)

type ActorDetails struct {
	ID   uuid.UUID
	Type db.GlobalActorType
}

func (a ActorDetails) IsSystem() bool {
	if a.Type == db.GlobalActorTypeSystem && a.ID == systemActorID {
		return true
	}

	return false
}
