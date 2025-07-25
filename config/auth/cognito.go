package auth

import (
	"github.com/Adgytec/adgytec-flow/utils/interfaces"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type authCognito struct {
	client *cognitoidentityprovider.Client
}

func (a *authCognito) CreateUser(username string) error {
	return nil
}

func (a *authCognito) DisableUser(username string) error {
	return nil
}

func (a *authCognito) EnableUser(username string) error {
	return nil
}

func CreateCognitoAuthClient(awsConfig aws.Config) interfaces.IAuth {
	return &authCognito{
		client: cognitoidentityprovider.NewFromConfig(awsConfig),
	}
}
