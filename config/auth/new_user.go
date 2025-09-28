package auth

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

func (a *authCognito) NewUser(ctx context.Context, username string) error {
	_, newUserErr := a.client.AdminCreateUser(
		ctx,
		&cognitoidentityprovider.AdminCreateUserInput{
			UserPoolId: aws.String(a.userPoolID),
			Username:   aws.String(username),
		},
	)
	if newUserErr != nil {
		var userExistsError *types.UsernameExistsException
		if errors.As(newUserErr, &userExistsError) {
			return &UserExistsError{
				username: username,
			}
		}

		return &AuthActionFailedError{
			username:   username,
			cause:      newUserErr,
			actionType: authActionTypeCreate,
		}
	}

	return nil
}
