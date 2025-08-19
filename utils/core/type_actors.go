package core

import "github.com/google/uuid"

type ActorType string

const (
	ActorTypeUser    ActorType = "user"
	ActorTypeApiKey  ActorType = "api-key"
	ActorTypeUnknown ActorType = "unknown"
)

func (a ActorType) Value() ActorType {
	switch a {
	case ActorTypeApiKey, ActorTypeUser:
		return a
	}

	return ActorTypeUnknown
}

type ActorDetails struct {
	ID   uuid.UUID
	Type ActorType
}
