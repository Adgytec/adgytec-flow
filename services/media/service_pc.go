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
}

func (pc *mediaServicePC) WithTransaction(db database.Database) MediaServicePC {
	serviceCopy := *pc.service
	serviceCopy.database = db

	return &mediaServicePC{
		service: &serviceCopy,
	}
}

func NewMediaServicePC(params mediaServiceParams) MediaServicePCTransaction {
	log.Printf("creating %s-service PC", serviceName)
	return &mediaServicePC{
		service: newMediaService(params),
	}
}
