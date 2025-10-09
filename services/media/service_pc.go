package media

import (
	"context"
	"log"

	"github.com/Adgytec/adgytec-flow/config/database"
)

type MediaServicePC interface {
	NewMediaItem(ctx context.Context, input NewMediaItemInfoWithStorageDetails) (*MediaUploadDetails, error)
	NewMediaItems(ctx context.Context, input []NewMediaItemInfoWithStorageDetails) ([]MediaUploadDetails, error)
}

type MediaServicePCTransaction interface {
	WithTransaction(db database.Database) MediaServicePC
}

type mediaServicePC struct {
	service *mediaService
	params  mediaServiceParams
}

func (pc *mediaServicePC) WithTransaction(db database.Database) MediaServicePC {
	return newMediaServiceActions(pc.params)
}

func NewMediaServicePC(params mediaServiceParams) MediaServicePCTransaction {
	log.Printf("creating %s-service PC", serviceName)
	return &mediaServicePC{
		service: newMediaService(params),
		params:  params,
	}
}
