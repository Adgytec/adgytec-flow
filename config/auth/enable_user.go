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
		var userExistsError *types.UsernameExistsException
		if errors.As(enableUserErr, &userExistsError) {
			return &UserExistsError{
				username: username,
			}
		}

		return &AuthActionFailedError{
			username:   username,
			cause:      enableUserErr,
			actionType: authActionEnableUser,
		}
	}

	return nil
}
