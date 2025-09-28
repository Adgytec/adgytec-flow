package auth

import "github.com/google/uuid"

func (a *authCognito) ValidateUserAccessToken(accessToken string) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}
