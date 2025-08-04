package auth

import (
	"log"
	"os"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type authCognito struct {
	client         *cognitoidentityprovider.Client
	userPoolId     string
	userPoolRegion string
}

func (a *authCognito) CreateUser(username string) (string, error) {
	return "", nil
}

func (a *authCognito) DisableUser(username string) error {
	return nil
}

func (a *authCognito) EnableUser(username string) error {
	return nil
}

func (a *authCognito) AddUserToManagement(username string) error {
	return nil
}

func CreateCognitoAuthClient(awsConfig aws.Config) core.IAuth {
	log.Println("init authentication cognito")
	return &authCognito{
		client:         cognitoidentityprovider.NewFromConfig(awsConfig),
		userPoolId:     os.Getenv("AWS_USER_POOL_ID"),
		userPoolRegion: os.Getenv("AWS_USER_POOL_REGION"),
	}
}
