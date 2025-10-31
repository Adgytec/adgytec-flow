package auth

import (
	"errors"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (a *authCognito) ValidateUserAccessToken(accessToken string) (uuid.UUID, error) {
	jwtToken, jwtParseErr := jwt.Parse(accessToken, a.jwtKeyfunc.Keyfunc)
	if jwtParseErr != nil {
		if errors.Is(jwtParseErr, jwt.ErrTokenMalformed) ||
			errors.Is(jwtParseErr, jwt.ErrTokenSignatureInvalid) ||
			errors.Is(jwtParseErr, jwt.ErrTokenExpired) ||
			errors.Is(jwtParseErr, jwt.ErrTokenNotValidYet) {
			return uuid.Nil, &InvalidAccessTokenError{
				cause: jwtParseErr,
			}
		}

		return uuid.Nil, &AuthActionFailedError{
			cause:      jwtParseErr,
			actionType: authActionTypeValidateAccessToken,
		}
	}

	invalidTokenErr := &InvalidAccessTokenError{
		cause: ErrInvalidAccessToken,
	}

	if !jwtToken.Valid {
		return uuid.Nil, invalidTokenErr
	}

	// get username from claims
	claims, claimsOK := jwtToken.Claims.(jwt.MapClaims)
	if !claimsOK {
		return uuid.Nil, invalidTokenErr
	}

	// cognito access token contains claim 'username' for user's username field
	username, usernameOK := claims["username"].(string)
	if !usernameOK {
		return uuid.Nil, invalidTokenErr
	}

	userID := core.GetUserIDFromUsername(username)
	return userID, nil
}
