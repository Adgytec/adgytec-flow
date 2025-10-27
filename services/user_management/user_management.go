package usermanagement

import (
	"github.com/Adgytec/adgytec-flow/config/cache"
	"github.com/Adgytec/adgytec-flow/config/database"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/services/user"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type serviceParams interface {
	Database() database.Database
	UserService() user.UserServicePC
	CacheClient() cache.CacheClient
	IAMService() iam.IAMServicePC
}

type serviceMuxParams interface {
	serviceParams
	Middleware() core.MiddlewarePC
}

type userManagementService struct {
	db          database.Database
	userService user.UserServicePC
	iam         iam.IAMServicePC
}

func newService(params serviceParams) *userManagementService {
	return &userManagementService{
		db:          params.Database(),
		userService: params.UserService(),
		iam:         params.IAMService(),
	}
}
