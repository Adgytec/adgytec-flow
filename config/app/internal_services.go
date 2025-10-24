package app

import (
	"github.com/Adgytec/adgytec-flow/services/appmiddleware"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/services/media"
	org "github.com/Adgytec/adgytec-flow/services/organization"
	"github.com/Adgytec/adgytec-flow/services/user"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type internalServices struct {
	iamService   iam.IAMServicePC
	userService  user.UserServicePC
	middleware   core.MiddlewarePC
	mediaService media.MediaServicePC
	orgService   org.OrgServicePC
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

func (s *internalServices) Organization() org.OrgServicePC {
	return s.orgService
}

func newInternalService(externalService appExternalServices) (appInternalServices, error) {
	internalService := internalServices{}
	appInstance := &app{
		appExternalServices: externalService,
		appInternalServices: &internalService,
	}

	// Initialize internal services. The order of initialization is important.
	internalService.iamService = iam.NewIAMServicePC(externalService)
	internalService.mediaService = media.NewMediaServicePC(appInstance)

	internalService.userService = user.NewUserServicePC(appInstance)
	internalService.orgService = org.NewOrgMux(appInstance)

	internalService.middleware = appmiddleware.NewAppMiddlewarePC(appInstance)

	return &internalService, nil
}
