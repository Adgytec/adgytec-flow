package auth

import (
	"log"
	"os"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/google/uuid"
)

type authCognito struct {
	client         *cognitoidentityprovider.Client
	userPoolID     string
	userPoolRegion string
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

func (a *authCognito) ValidateUserAccessToken(accessToken string) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

func (a *authCognito) ValidateAPIKey(apiKey string) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

func NewCognitoAuthClient(awsConfig aws.Config) core.IAuth {
	log.Println("init authentication cognito")
	return &authCognito{
		client:         cognitoidentityprovider.NewFromConfig(awsConfig),
		userPoolID:     os.Getenv("AWS_USER_POOL_ID"),
		userPoolRegion: os.Getenv("AWS_USER_POOL_REGION"),
	}
}
