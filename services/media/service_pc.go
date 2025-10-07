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

// this is created inorder to avoid calling multiple transaction inside a transaction
type MediaServicePCWithTransaction interface {
	MediaServicePC
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

func NewMediaServicePC(params mediaServiceParams) MediaServicePCWithTransaction {
	log.Printf("creating %s-service PC", serviceName)
	return &mediaServicePC{
		service: newMediaService(params),
	}
}
