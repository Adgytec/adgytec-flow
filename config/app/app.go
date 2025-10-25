package app

import (
	"context"

	"github.com/Adgytec/adgytec-flow/config/auth"
	"github.com/Adgytec/adgytec-flow/config/cache"
	"github.com/Adgytec/adgytec-flow/config/cdn"
	"github.com/Adgytec/adgytec-flow/config/communication"
	"github.com/Adgytec/adgytec-flow/config/database"
	"github.com/Adgytec/adgytec-flow/config/storage"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/services/media"
	org "github.com/Adgytec/adgytec-flow/services/organization"
	"github.com/Adgytec/adgytec-flow/services/user"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/staging"
)

type appExternalServices interface {
	Auth() auth.Auth
	Database() database.Database
	Communication() communication.Communication
	Storage() storage.Storage
	CDN() cdn.CDN
	Shutdown(ctx context.Context)
	CacheClient() cache.CacheClient
}

type appInternalServices interface {
	IAMService() iam.IAMServicePC
	UserService() user.UserServicePC
	Middleware() core.MiddlewarePC
	MediaWithTransaction() media.MediaServicePC
	Organization() org.OrgServicePC
}

type App interface {
	appExternalServices
	appInternalServices

	AddServices(details []staging.Details)
	Services() staging.Services
}

type app struct {
	appExternalServices
	appInternalServices

	services staging.Services
}

func (a *app) Services() staging.Services {
	return a.services
}

func (a *app) AddServices(details []staging.Details) {
	a.services = staging.NewServices(details)
}
