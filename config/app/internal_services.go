package app

import (
	"github.com/Adgytec/adgytec-flow/services/access_management"
	app_middleware "github.com/Adgytec/adgytec-flow/services/middleware"
	"github.com/Adgytec/adgytec-flow/services/user"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type internalServices struct {
	accessManagement core.AccessManagementPC
	userService      core.UserServicePC
	middleware       core.MiddlewarePC
}

func (s *internalServices) AccessManagement() core.AccessManagementPC {
	return s.accessManagement
}

func (s *internalServices) UserService() core.UserServicePC {
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
	internalService.accessManagement = access_management.NewAccessManagementPC(externalService)
	internalService.userService = user.NewUserServicePC(appInstance)
	internalService.middleware = app_middleware.NewAppMiddlewarePC(appInstance)

	return &internalService
}
