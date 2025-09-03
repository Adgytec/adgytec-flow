package access_management

import (
	"github.com/Adgytec/adgytec-flow/config/cache"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type iAccessManagementParams interface {
	Database() core.IDatabase
	CacheClient() core.ICacheClient
}

type iAccessManagementMuxParams interface {
	iAccessManagementParams
	Middleware() core.IMiddlewarePC
}

type accessManagement struct {
	db              core.IDatabase
	permissionCache core.ICache[bool]
}

func newAccessManagementService(params iAccessManagementParams) *accessManagement {
	return &accessManagement{
		db:              params.Database(),
		permissionCache: cache.NewCache[bool](params.CacheClient(), "access-management"),
	}
}
