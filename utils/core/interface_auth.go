package core

import "github.com/google/uuid"

type IAuth interface {
	CreateUser(string) error
	DisableUser(string) error
	EnableUser(string) error
	ValidateUserAccessToken(string) (uuid.UUID, error)
}
