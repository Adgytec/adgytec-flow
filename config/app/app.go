package app

import (
	"github.com/Adgytec/adgytec-flow/services/user"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type appExternalServices interface {
	Auth() core.Auth
	Database() core.Database
	Communication() core.Communication
	Storage() core.Storage
	CDN() core.CDN
	Shutdown()
	CacheClient() core.CacheClient
}

type appInternalServices interface {
	AccessManagement() core.AccessManagementPC
	UserService() user.PC
	Middleware() core.MiddlewarePC
}

type App interface {
	appExternalServices
	appInternalServices
}

type app struct {
	appExternalServices
	appInternalServices
}
