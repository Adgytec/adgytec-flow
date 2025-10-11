package auth

import (
	"fmt"
	"net/url"
	"os"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/rs/zerolog/log"
)

type authCognito struct {
	authCommon
	client         *cognitoidentityprovider.Client
	userPoolID     string
	userPoolRegion string
	jwtKeyfunc     keyfunc.Keyfunc
}

func NewCognitoAuthClient(awsConfig aws.Config, apiURL *url.URL) (Auth, error) {
	log.Info().Msg("new cognito auth client")
	userPoolID := os.Getenv("AWS_USER_POOL_ID")
	userPoolRegion := os.Getenv("AWS_USER_POOL_REGION")

	if userPoolID == "" || userPoolRegion == "" {
		return nil, ErrInvalidAuthConfig
	}

	jwkSetEndpoint := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", userPoolRegion, userPoolID)

	jwtKeyfunc, keyFuncErr := keyfunc.NewDefault([]string{jwkSetEndpoint})
	if keyFuncErr != nil {
		return nil, &JwtKeyFuncError{
			cause: keyFuncErr,
		}
	}

	authCommon, authCommonErr := newAuthCommon(apiURL)
	if authCommonErr != nil {
		return nil, authCommonErr
	}

	return &authCognito{
		authCommon:     *authCommon,
		client:         cognitoidentityprovider.NewFromConfig(awsConfig),
		userPoolID:     userPoolID,
		userPoolRegion: userPoolRegion,
		jwtKeyfunc:     jwtKeyfunc,
	}, nil
}
