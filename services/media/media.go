package media

import "github.com/Adgytec/adgytec-flow/config/storage"

type mediaServiceParams interface {
	Storage() storage.Storage
}

type mediaService struct {
	storage storage.Storage
}

func newMediaService(params mediaServiceParams) *mediaService {
	return &mediaService{
		storage: params.Storage(),
	}
}
