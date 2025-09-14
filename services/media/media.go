package media

import (
	"github.com/Adgytec/adgytec-flow/config/database"
	"github.com/Adgytec/adgytec-flow/config/storage"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type mediaServiceParams interface {
	Storage() storage.Storage
	Database() database.Database
}

type mediaServiceMuxParams interface {
	mediaServiceParams
	Middleware() core.MiddlewarePC
}

type mediaService struct {
	storage storage.Storage
}

func newMediaService(params mediaServiceParams) *mediaService {
	return &mediaService{
		storage: params.Storage(),
	}
}
