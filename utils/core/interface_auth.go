package core

import "github.com/google/uuid"

type IAuth interface {
	CreateUser(string) error
	DisableUser(string) error
	EnableUser(string) error
	ValidateUserAccessToken(string) (uuid.UUID, error)

	// this only checks if the API key is in required format as described in the application doc
	// further validation like if this api key actually exists is done later on
	ValidateAPIKey(string) (uuid.UUID, error)
}
