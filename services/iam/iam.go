package iam

import (
	"github.com/Adgytec/adgytec-flow/config/cache"
	"github.com/Adgytec/adgytec-flow/config/serializer"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type iamServiceParams interface {
	Database() core.Database
	CacheClient() core.CacheClient
}

type iamServiceMuxParams interface {
	iamServiceParams
	Middleware() core.MiddlewarePC
}

type iamService struct {
	db              core.Database
	permissionCache core.Cache[bool]
}

func newIAMService(params iamServiceParams) *iamService {
	return &iamService{
		db:              params.Database(),
		permissionCache: cache.NewCache[bool](params.CacheClient(), serializer.NewJSONSerializer[bool](), "iam"),
	}
}
