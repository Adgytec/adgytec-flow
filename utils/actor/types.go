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

type ActorDetails struct {
	ID   uuid.UUID
	Type db.GlobalActorType
}
