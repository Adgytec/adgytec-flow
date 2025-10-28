package auth

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

func (a *authCognito) DisableUser(ctx context.Context, username string) error {
	_, disableUserErr := a.client.AdminDisableUser(ctx,
		&cognitoidentityprovider.AdminDisableUserInput{
			UserPoolId: aws.String(a.userPoolID),
			Username:   aws.String(username),
		},
	)

	if disableUserErr != nil {
		var userExistsError *types.UsernameExistsException
		if errors.As(disableUserErr, &userExistsError) {
			return &UserExistsError{
				username: username,
			}
		}

		return &AuthActionFailedError{
			username:   username,
			cause:      disableUserErr,
			actionType: authActionDisableUser,
		}
	}

	return nil
}
