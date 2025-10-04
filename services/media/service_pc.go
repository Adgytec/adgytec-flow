package media

import (
	"context"
	"log"

	"github.com/Adgytec/adgytec-flow/config/database"
)

// Contains methods that can be run without a transaction.
type MediaServicePC interface {
	NewMediaItem(ctx context.Context, input NewMediaItemInputWithBucketPrefix) (*NewMediaItemOutput, error)
	NewMediaItems(ctx context.Context, input []NewMediaItemInputWithBucketPrefix) ([]NewMediaItemOutput, error)
}

// The main service interface that provides access to both.
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
