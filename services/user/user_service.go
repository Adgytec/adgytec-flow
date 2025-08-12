package user

import (
	"github.com/Adgytec/adgytec-flow/config/cache"
	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type iUserServiceParams interface {
	Database() core.IDatabase
	Auth() core.IAuth
	AccessManagement() core.IAccessManagementPC
	CDN() core.ICDN
	CacheClient() core.ICacheClient
}

type userService struct {
	db                    core.IDatabase
	auth                  core.IAuth
	accessManagement      core.IAccessManagementPC
	cdn                   core.ICDN
	getUserCache          core.ICache[db_actions.GlobalUser]
	getUserListCache      core.ICache[[]db_actions.GlobalUser]
	userLastAccessedCache core.ICache[bool]
}

func createUserService(params iUserServiceParams) *userService {
	return &userService{
		db:                    params.Database(),
		auth:                  params.Auth(),
		accessManagement:      params.AccessManagement(),
		cdn:                   params.CDN(),
		getUserCache:          cache.CreateNewCache[db_actions.GlobalUser](params.CacheClient(), "user"),
		getUserListCache:      cache.CreateNewCache[[]db_actions.GlobalUser](params.CacheClient(), "user-list"),
		userLastAccessedCache: cache.CreateNewCache[bool](params.CacheClient(), "user-last-accessed"),
	}
}
