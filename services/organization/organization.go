package org

import (
	"github.com/Adgytec/adgytec-flow/config/cache"
	"github.com/Adgytec/adgytec-flow/config/database"
	"github.com/Adgytec/adgytec-flow/config/serializer"
	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/services/user"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type orgServiceParams interface {
	Database() database.Database
	User() user.UserServicePC
	CacheClient() cache.CacheClient
}

type orgServiceMuxParams interface {
	orgServiceParams
	Middleware() core.MiddlewarePC
}

type orgService struct {
	db                           database.Database
	userService                  user.UserServicePC
	coreServiceRestrictionsCache cache.Cache[[]db.GetCoreServiceRestrictionsRow]
}

func newOrgService(params orgServiceParams) *orgService {
	return &orgService{
		db:          params.Database(),
		userService: params.User(),
		coreServiceRestrictionsCache: cache.NewCache(
			params.CacheClient(),
			serializer.NewGobSerializer[[]db.GetCoreServiceRestrictionsRow](),
			"core-service-restriction",
		),
	}
}
