package user

import (
	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type iUserServiceParams interface {
	Database() core.IDatabase
	Auth() core.IAuth
	AccessManagement() core.IAccessManagementPC
	CacheClient() core.ICacheClient
}

type userService struct {
	db                    core.IDatabase
	auth                  core.IAuth
	accessManagement      core.IAccessManagementPC
	getUserCache          core.ICache[db_actions.GlobalUser]
	getUserListCache      core.ICache[[]db_actions.GlobalUser]
	userLastAccessedCache core.ICache[bool]
}
