package org

import (
	"github.com/Adgytec/adgytec-flow/config/cache"
	"github.com/Adgytec/adgytec-flow/config/database"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/services/media"
	"github.com/Adgytec/adgytec-flow/services/user"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/staging"
)

type orgServiceParams interface {
	Database() database.Database
	UserService() user.UserServicePC
	CacheClient() cache.CacheClient
	MediaWithTransaction() media.MediaServicePC
	IAMService() iam.IAMServicePC
	Services() staging.Services
}

type orgServiceMuxParams interface {
	orgServiceParams
	Middleware() core.MiddlewarePC
}

type orgService struct {
	db             database.Database
	userService    user.UserServicePC
	media          media.MediaServicePC
	iam            iam.IAMServicePC
	serviceDetails staging.Services
}

func newOrgService(params orgServiceParams) *orgService {
	return &orgService{
		db:             params.Database(),
		userService:    params.UserService(),
		media:          params.MediaWithTransaction(),
		iam:            params.IAMService(),
		serviceDetails: params.Services(),
	}
}
