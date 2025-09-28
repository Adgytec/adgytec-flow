package auth

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type authCognito struct {
	authCommon
	client         *cognitoidentityprovider.Client
	userPoolID     string
	userPoolRegion string
}

func NewCognitoAuthClient(awsConfig aws.Config) Auth {
	log.Println("init authentication cognito")
	return &authCognito{
		authCommon:     newAuthCommon(),
		client:         cognitoidentityprovider.NewFromConfig(awsConfig),
		userPoolID:     os.Getenv("AWS_USER_POOL_ID"),
		userPoolRegion: os.Getenv("AWS_USER_POOL_REGION"),
	}
}
