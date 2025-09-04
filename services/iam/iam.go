package iam

import (
	"github.com/Adgytec/adgytec-flow/config/cache"
	"github.com/Adgytec/adgytec-flow/config/serializer"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type iamParams interface {
	Database() core.Database
	CacheClient() core.CacheClient
}

type iamServiceMuxParams interface {
	iamParams
	Middleware() core.MiddlewarePC
}

type iam struct {
	db              core.Database
	permissionCache core.Cache[bool]
}

func newIAMService(params iamParams) *iam {
	return &iam{
		db:              params.Database(),
		permissionCache: cache.NewCache[bool](params.CacheClient(), serializer.NewJSONSerializer[bool](), "iam"),
	}
}
