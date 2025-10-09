package media

import (
	"net/url"

	"github.com/Adgytec/adgytec-flow/config/auth"
	"github.com/Adgytec/adgytec-flow/config/database"
	"github.com/Adgytec/adgytec-flow/config/storage"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

const zeroMime string = "application/octet-stream"

type mediaServiceParams interface {
	Storage() storage.Storage
	Database() database.Database
	Auth() auth.Auth
}

type mediaServiceMuxParams interface {
	mediaServiceParams
	Middleware() core.MiddlewarePC
}

type mediaService struct {
	storage  storage.Storage
	database database.Database
	auth     auth.Auth
	apiURL   url.URL
}

func newMediaService(params mediaServiceParams) *mediaService {
	return &mediaService{
		storage:  params.Storage(),
		database: params.Database(),
		auth:     params.Auth(),
	}
}
