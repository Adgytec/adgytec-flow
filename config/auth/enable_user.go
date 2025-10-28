package auth

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

func (a *authCognito) EnableUser(ctx context.Context, username string) error {
	_, enableUserErr := a.client.AdminEnableUser(ctx,
		&cognitoidentityprovider.AdminEnableUserInput{
			UserPoolId: aws.String(a.userPoolID),
			Username:   aws.String(username),
		},
	)

	if enableUserErr != nil {
		var userNotFoundError *types.UserNotFoundException
		if errors.As(enableUserErr, &userNotFoundError) {
			return &UserNotFoundError{
				username: username,
			}
		}

		return &AuthActionFailedError{
			username:   username,
			cause:      enableUserErr,
			actionType: authActionTypeEnableUser,
		}
	}

	return nil
}
