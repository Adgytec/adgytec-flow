package app

import (
	"github.com/Adgytec/adgytec-flow/services/iam"
	app_middleware "github.com/Adgytec/adgytec-flow/services/middleware"
	"github.com/Adgytec/adgytec-flow/services/user"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type internalServices struct {
	iamService  iam.IAMServicePC
	userService user.UserServicePC
	middleware  core.MiddlewarePC
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

func newInternalService(externalService appExternalServices) appInternalServices {
	internalService := internalServices{}
	appInstance := &app{
		appExternalServices: externalService,
		appInternalServices: &internalService,
	}

	// Initialize internal services. The order of initialization is important.
	internalService.iamService = iam.NewIAMServicePC(externalService)
	internalService.userService = user.NewUserServicePC(appInstance)
	internalService.middleware = app_middleware.NewAppMiddlewarePC(appInstance)

	return &internalService
}
