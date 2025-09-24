package media

import (
	"github.com/Adgytec/adgytec-flow/config/database"
	"github.com/Adgytec/adgytec-flow/config/storage"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

// 16 mega byte
const singlepartUploadLimit int64 = 16 * (1 << 20)

// 1 giga byte
const multipartUploadLimit int64 = 1 * (1 << 30)

// 5 mega byte
const multipartPartSize int64 = 5 * (1 << 20)

const mediaUploadLimit int = 50

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
