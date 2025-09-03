package access_management

import (
	"github.com/Adgytec/adgytec-flow/config/cache"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type accessManagementParams interface {
	Database() core.Database
	CacheClient() core.CacheClient
}

type accessManagementMuxParams interface {
	accessManagementParams
	Middleware() core.MiddlewarePC
}

type accessManagement struct {
	db              core.Database
	permissionCache core.Cache[bool]
}

func newAccessManagementService(params accessManagementParams) *accessManagement {
	return &accessManagement{
		db:              params.Database(),
		permissionCache: cache.NewCache[bool](params.CacheClient(), "access-management"),
	}
}
