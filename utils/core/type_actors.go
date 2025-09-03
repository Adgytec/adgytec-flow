package core

import (
	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/google/uuid"
)

type ActorDetails struct {
	ID   uuid.UUID
	Type db.GlobalActorType
}
