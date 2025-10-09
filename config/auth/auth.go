package auth

import (
	"context"
	"net/url"
	"os"
	"time"

	"github.com/google/uuid"
)

type Auth interface {
	NewUser(ctx context.Context, username string) error
	ValidateUserAccessToken(accessToken string) (uuid.UUID, error)

	// this only checks if the API key is in required format as described in the application doc
	// further validation like if this api key actually exists is done later on
	ValidateAPIKey(apiKey string) (uuid.UUID, error)
	NewSignedURL(path string, query map[string]string, expireAfter time.Duration) (*url.URL, error)
	NewSignedURLWithActor(ctx context.Context, path string, query map[string]string, expireAfter time.Duration) (*url.URL, error)
	ValidateSignedURL(signedURL url.URL) error
	ValidateSignedURLWithActor(ctx context.Context, signedURL url.URL) error
}

// authCommon contains method impl that are independent of external authentication provider
type authCommon struct {
	secret []byte
	apiURL *url.URL
}

func newAuthCommon(apiURL *url.URL) (*authCommon, error) {
	hmacSecret := os.Getenv("HMAC_SECRET")
	if hmacSecret == "" {
		return nil, ErrInvalidHMACSecret
	}

	return &authCommon{
		secret: []byte(hmacSecret),
		apiURL: apiURL,
	}, nil
}

// used in auth errors
type authActionType string

const (
	authActionTypeCreate              authActionType = "auth-user-create"
	authActionTypeValidateAccessToken authActionType = "auth-validate-user-access-token"
)
