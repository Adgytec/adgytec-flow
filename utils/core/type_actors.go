package core

import (
	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/google/uuid"
)

// type ActorType string

// const (
// 	ActorTypeUser    ActorType = "user"
// 	ActorTypeApiKey  ActorType = "api_key"
// 	ActorTypeUnknown ActorType = "unknown"
// )
//
// func (a ActorType) Value() ActorType {
// 	switch a {
// 	case ActorTypeApiKey, ActorTypeUser:
// 		return a
// 	}
//
// 	return ActorTypeUnknown
// }

type ActorDetails struct {
	ID   uuid.UUID
	Type db_actions.GlobalActorType
}
