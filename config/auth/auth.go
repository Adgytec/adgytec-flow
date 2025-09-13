package auth

import "github.com/google/uuid"

type Auth interface {
	NewUser(username string) error
	DisableUser(username string) error
	EnableUser(username string) error
	ValidateUserAccessToken(accessToken string) (uuid.UUID, error)

	// this only checks if the API key is in required format as described in the application doc
	// further validation like if this api key actually exists is done later on
	ValidateAPIKey(apiKey string) (uuid.UUID, error)
	NewSignedHash(payload ...[]byte) string
}

// authCommon contains method impl that are independent of external authentication provider
type authCommon struct{}

func (a *authCommon) ValidateAPIKey(apiKey string) (uuid.UUID, error) {
	return uuid.Nil, nil
}

func (a *authCommon) NewSignedHash(payload ...[]byte) string {
	return ""
}

func newAuthCommon() authCommon {
	return authCommon{}
}

// used in auth errors
type authActionType string

const (
	authActionTypeCreate authActionType = "auth-user-create"
)
