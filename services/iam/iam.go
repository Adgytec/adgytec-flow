package iam

import (
	"github.com/Adgytec/adgytec-flow/config/cache"
	"github.com/Adgytec/adgytec-flow/config/database"
	"github.com/Adgytec/adgytec-flow/config/serializer"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type iamServiceParams interface {
	Database() database.Database
	CacheClient() cache.CacheClient
}

type iamServiceMuxParams interface {
	iamServiceParams
	Middleware() core.MiddlewarePC
}

type iamService struct {
	db              database.Database
	permissionCache cache.Cache[bool]
}

func newIAMService(params iamServiceParams) *iamService {
	return &iamService{
		db:              params.Database(),
		permissionCache: cache.NewCache[bool](params.CacheClient(), serializer.NewJSONSerializer[bool](), "iam"),
	}
}
