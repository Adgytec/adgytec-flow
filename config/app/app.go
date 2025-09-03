package app

import "github.com/Adgytec/adgytec-flow/utils/core"

type appExternalServices interface {
	Auth() core.Auth
	Database() core.Database
	Communication() core.Communicaiton
	Storage() core.Storage
	CDN() core.CDN
	Shutdown()
	CacheClient() core.CacheClient
}

type appInternalServices interface {
	AccessManagement() core.AccessManagementPC
	UserService() core.UserServicePC
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
