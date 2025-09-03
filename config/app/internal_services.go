package app

import (
	"github.com/Adgytec/adgytec-flow/services/access_management"
	app_middleware "github.com/Adgytec/adgytec-flow/services/middleware"
	"github.com/Adgytec/adgytec-flow/services/user"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type internalServices struct {
	accessManagement core.IAccessManagementPC
	userService      core.IUserServicePC
	middleware       core.IMiddlewarePC
}

func (s *internalServices) AccessManagement() core.IAccessManagementPC {
	return s.accessManagement
}

func (s *internalServices) UserService() core.IUserServicePC {
	return s.userService
}

func (s *internalServices) Middleware() core.IMiddlewarePC {
	return s.middleware
}

func createInternalService(externalService iAppExternalServices) iAppInternalServices {
	internalService := internalServices{}
	appInstance := &app{
		iAppExternalServices: externalService,
		iAppInternalServices: &internalService,
	}

	// Initialize internal services. The order of initialization is important.
	internalService.accessManagement = access_management.NewAccessManagementPC(externalService)
	internalService.userService = user.NewUserServicePC(appInstance)
	internalService.middleware = app_middleware.CreateAppMiddlewarePC(appInstance)

	return &internalService
}
