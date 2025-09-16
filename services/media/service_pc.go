package media

import (
	"context"
	"log"

	"github.com/Adgytec/adgytec-flow/config/database"
	"github.com/google/uuid"
)

// Contains methods that can be run without a transaction.
type MediaServicePC interface {
	NewMediaItem(ctx context.Context, input NewMediaItemInput) (NewMediaItemOutput, error)
	NewMediaItems(ctx context.Context, input []NewMediaItemInput) ([]NewMediaItemOutput, error)
}

// Contains methods that MUST be run within a transaction.
type TransactionalMediaServicePC interface {
	CompleteMediaItemUpload(ctx context.Context, mediaID uuid.UUID) error
	CompleteMediaItemsUpload(ctx context.Context, mediaIDs []uuid.UUID) error
}

// The main service interface that provides access to both.
type MediaServicePCWithTransaction interface {
	MediaServicePC
	WithTransaction(db database.Database) TransactionalMediaServicePC
}

type mediaServicePC struct {
	service *mediaService
}

func (pc *mediaServicePC) WithTransaction(db database.Database) TransactionalMediaServicePC {
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
