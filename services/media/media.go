package media

import (
	"github.com/Adgytec/adgytec-flow/config/database"
	"github.com/Adgytec/adgytec-flow/config/storage"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

// 16 mega byte
const singlepartUploadLimit = 16 * (1 << 20)

type mediaServiceParams interface {
	Storage() storage.Storage
	Database() database.Database
}

type mediaServiceMuxParams interface {
	mediaServiceParams
	Middleware() core.MiddlewarePC
}

type mediaService struct {
	storage  storage.Storage
	database database.Database
}

func newMediaService(params mediaServiceParams) *mediaService {
	return &mediaService{
		storage:  params.Storage(),
		database: params.Database(),
	}
}
