package app

import (
	"github.com/Adgytec/adgytec-flow/config/auth"
	"github.com/Adgytec/adgytec-flow/config/cache"
	"github.com/Adgytec/adgytec-flow/config/cdn"
	"github.com/Adgytec/adgytec-flow/config/communication"
	"github.com/Adgytec/adgytec-flow/config/storage"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/services/user"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type appExternalServices interface {
	Auth() auth.Auth
	Database() core.Database
	Communication() communication.Communication
	Storage() storage.Storage
	CDN() cdn.CDN
	Shutdown()
	CacheClient() cache.CacheClient
}

type appInternalServices interface {
	IAMService() iam.IAMServicePC
	UserService() user.UserServicePC
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
