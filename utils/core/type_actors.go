package core

import "github.com/google/uuid"

type ActorType string

const (
	ActorTypeUser    ActorType = "user"
	ActorTypeApiKey  ActorType = "api-key"
	ActorTypeInvalid ActorType = "invalid"
)

func (a ActorType) GetPermissionEntityType() PermissionEntityType {
	switch a {
	case ActorTypeUser:
		return PermissionEntityTypeUser
	case ActorTypeApiKey:
		return PermissionEntityTypeAPIKey
	}

	return PermissionEntityTypeUser
}

func (a ActorType) Value() ActorType {
	switch a {
	case ActorTypeApiKey, ActorTypeUser:
		return a
	}

	return ActorTypeInvalid
}

type ActorDetials struct {
	ID   uuid.UUID
	Type ActorType
}
