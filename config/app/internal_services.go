package app

import (
	"errors"
	"net/url"
	"os"
	"strings"

	"github.com/Adgytec/adgytec-flow/services/appmiddleware"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/services/media"
	"github.com/Adgytec/adgytec-flow/services/user"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type internalServices struct {
	iamService   iam.IAMServicePC
	userService  user.UserServicePC
	middleware   core.MiddlewarePC
	mediaService media.MediaServicePC
	apiURL       *url.URL
}

func (s *internalServices) IAMService() iam.IAMServicePC {
	return s.iamService
}

func (s *internalServices) UserService() user.UserServicePC {
	return s.userService
}

func (s *internalServices) Middleware() core.MiddlewarePC {
	return s.middleware
}

func (s *internalServices) MediaWithTransaction() media.MediaServicePC {
	return s.mediaService
}

func (s *internalServices) ApiURL() *url.URL {
	return s.apiURL
}

func newInternalService(externalService appExternalServices) (appInternalServices, error) {
	// parse api endpoint
	urlString := os.Getenv("API_ENDPOINT")
	if strings.TrimSpace(urlString) == "" {
		return nil, errors.New("missing API_ENDPOINT env variable")
	}

	apiURL, parseErr := url.Parse(urlString)
	if parseErr != nil {
		return nil, parseErr
	}

	internalService := internalServices{
		apiURL: apiURL,
	}
	appInstance := &app{
		appExternalServices: externalService,
		appInternalServices: &internalService,
	}

	// Initialize internal services. The order of initialization is important.
	internalService.iamService = iam.NewIAMServicePC(externalService)
	internalService.mediaService = media.NewMediaServicePC(externalService)

	internalService.userService = user.NewUserServicePC(appInstance)
	internalService.middleware = appmiddleware.NewAppMiddlewarePC(appInstance)

	return &internalService, nil
}
