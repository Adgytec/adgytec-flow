package app

import (
	"github.com/Adgytec/adgytec-flow/services/access_management"
	"github.com/Adgytec/adgytec-flow/services/user"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type internalServices struct {
	accessManagement core.IAccessManagementPC
	userService      core.IUserServicePC
}

func (s *internalServices) AccessManagement() core.IAccessManagementPC {
	return s.accessManagement
}

func (s *internalServices) UserService() core.IUserServicePC {
	return s.userService
}

func createInternalService(externalService iAppExternalServices) iAppInternalServices {
	internalService := internalServices{}
	appInstance := &app{
		iAppExternalServices: externalService,
		iAppInternalServices: &internalService,
	}

	// Initialize internal services. The order of initialization is important.
	internalService.accessManagement = access_management.CreateAccessManagementPC(externalService)
	internalService.userService = user.CreateUserServicePC(appInstance)

	return &internalService
}
