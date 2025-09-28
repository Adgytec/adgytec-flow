package auth

import (
	"fmt"
	"log"
	"os"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type authCognito struct {
	authCommon
	client         *cognitoidentityprovider.Client
	userPoolID     string
	userPoolRegion string
	jwtKeyfunc     keyfunc.Keyfunc
}

func NewCognitoAuthClient(awsConfig aws.Config) Auth {
	log.Println("init authentication cognito")

	userPoolID := os.Getenv("AWS_USER_POOL_ID")
	userPoolRegion := os.Getenv("AWS_USER_POOL_REGION")

	if userPoolID == "" || userPoolRegion == "" {
		log.Fatal("can't find env values for AWS_USER_POOL_ID or AWS_USER_POOL_REGION")
	}

	jwkSetEndpoint := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", userPoolRegion, userPoolID)

	jwtKeyfunc, keyFuncErr := keyfunc.NewDefault([]string{jwkSetEndpoint})
	if keyFuncErr != nil {
		log.Fatalf("Failed to create a keyfunc.Keyfunc from the server's URL.\nError: %s", keyFuncErr)
	}
	return &authCognito{
		authCommon:     newAuthCommon(),
		client:         cognitoidentityprovider.NewFromConfig(awsConfig),
		userPoolID:     userPoolID,
		userPoolRegion: userPoolRegion,
		jwtKeyfunc:     jwtKeyfunc,
	}
}
