package auth

import (
	"context"
	"os"

	"github.com/google/uuid"
)

type Auth interface {
	NewUser(ctx context.Context, username string) error
	ValidateUserAccessToken(accessToken string) (uuid.UUID, error)

	// this only checks if the API key is in required format as described in the application doc
	// further validation like if this api key actually exists is done later on
	ValidateAPIKey(apiKey string) (uuid.UUID, error)
	NewSignedHash(payload ...[]byte) (string, error)
	CompareSignedHash(hash string, payload ...[]byte) error
}

// authCommon contains method impl that are independent of external authentication provider
type authCommon struct {
	secret []byte
}

func newAuthCommon() (*authCommon, error) {
	hmacSecret := os.Getenv("HMAC_SECRET")
	if hmacSecret == "" {
		return nil, ErrInvalidHMACSecret
	}

	return &authCommon{
		secret: []byte(hmacSecret),
	}, nil
}

// used in auth errors
type authActionType string

const (
	authActionTypeCreate              authActionType = "auth-user-create"
	authActionTypeValidateAccessToken authActionType = "auth-validate-user-access-token"
)
