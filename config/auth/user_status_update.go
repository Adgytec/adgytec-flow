package auth

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

func (a *authCognito) handleUserAction(err error, username string, actionType authActionType) error {
	if err == nil {
		return nil
	}

	var userNotFoundError *types.UserNotFoundException
	if errors.As(err, &userNotFoundError) {
		return &UserNotFoundError{username: username}
	}

	return &AuthActionFailedError{
		username:   username,
		cause:      err,
		actionType: actionType,
	}
}

func (a *authCognito) EnableUser(ctx context.Context, username string) error {
	_, enableUserErr := a.client.AdminEnableUser(ctx,
		&cognitoidentityprovider.AdminEnableUserInput{
			UserPoolId: aws.String(a.userPoolID),
			Username:   aws.String(username),
		},
	)

	return a.handleUserAction(enableUserErr, username, authActionTypeEnableUser)
}

func (a *authCognito) DisableUser(ctx context.Context, username string) error {
	_, disableUserErr := a.client.AdminDisableUser(ctx,
		&cognitoidentityprovider.AdminDisableUserInput{
			UserPoolId: aws.String(a.userPoolID),
			Username:   aws.String(username),
		},
	)

	return a.handleUserAction(disableUserErr, username, authActionTypeDisableUser)
}
