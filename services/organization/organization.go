package org

import (
	"github.com/Adgytec/adgytec-flow/config/cache"
	"github.com/Adgytec/adgytec-flow/config/database"
	"github.com/Adgytec/adgytec-flow/config/serializer"
	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/services/media"
	"github.com/Adgytec/adgytec-flow/services/user"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type orgServiceParams interface {
	Database() database.Database
	UserService() user.UserServicePC
	CacheClient() cache.CacheClient
	MediaWithTransaction() media.MediaServicePC
}

type orgServiceMuxParams interface {
	orgServiceParams
	Middleware() core.MiddlewarePC
}

type orgService struct {
	db                           database.Database
	userService                  user.UserServicePC
	media                        media.MediaServicePC
	coreServiceRestrictionsCache cache.Cache[[]db.GetCoreServiceRestrictionsRow]
}

func newOrgService(params orgServiceParams) *orgService {
	return &orgService{
		db:          params.Database(),
		userService: params.UserService(),
		media:       params.MediaWithTransaction(),
		coreServiceRestrictionsCache: cache.NewCache(
			params.CacheClient(),
			serializer.NewGobSerializer[[]db.GetCoreServiceRestrictionsRow](),
			"core-service-restriction",
		),
	}
}
