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
			reason:     jwtParseErr.Error(),
			actionType: authActionTypeValidateAccessToken,
		}
	}

	invalidTokenErr := &AuthActionFailedError{
		cause:      ErrInvalidAccessToken,
		reason:     ErrInvalidAccessToken.Error(),
		actionType: authActionTypeValidateAccessToken,
	}

	if !jwtToken.Valid {
		return uuid.Nil, invalidTokenErr
	}

	// get username from claims
	claims, claimsOK := jwtToken.Claims.(jwt.MapClaims)
	if !claimsOK {
		return uuid.Nil, invalidTokenErr
	}

	username, usernameOK := claims["username"].(string)
	if !usernameOK {
		return uuid.Nil, invalidTokenErr
	}

	userID := core.GetIDFromPayload([]byte(username))
	return userID, nil
}
