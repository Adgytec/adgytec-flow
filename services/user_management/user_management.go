package usermanagement

import (
	"github.com/Adgytec/adgytec-flow/config/cache"
	"github.com/Adgytec/adgytec-flow/config/database"
	"github.com/Adgytec/adgytec-flow/config/serializer"
	"github.com/Adgytec/adgytec-flow/database/models"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/services/user"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/pagination"
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
	db                 database.Database
	userService        user.UserServicePC
	iam                iam.IAMServicePC
	getUserListCache   cache.Cache[pagination.ResponsePagination[models.GlobalUser]]
	userGroupListCache cache.Cache[pagination.ResponsePagination[models.UserGroup]]
}

func newService(params serviceParams) *userManagementService {
	return &userManagementService{
		db:                 params.Database(),
		userService:        params.UserService(),
		iam:                params.IAMService(),
		getUserListCache:   cache.NewCache(params.CacheClient(), serializer.NewGobSerializer[pagination.ResponsePagination[models.GlobalUser]](), "management-user-list"),
		userGroupListCache: cache.NewCache(params.CacheClient(), serializer.NewGobSerializer[pagination.ResponsePagination[models.UserGroup]](), "user-group-list"),
	}
}
