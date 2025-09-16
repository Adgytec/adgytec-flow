package media

import (
	"context"
	"log"

	"github.com/Adgytec/adgytec-flow/config/database"
	"github.com/google/uuid"
)

type MediaServicePC interface {
	NewMediaItem(ctx context.Context, input NewMediaItemInput) (NewMediaItemOutput, error)
	NewMediaItems(ctx context.Context, input []NewMediaItemInput) ([]NewMediaItemOutput, error)
	CompleteMediaItemUpload(ctx context.Context, mediaID uuid.UUID) error
	CompleteMediaItemsUpload(ctx context.Context, mediaIDs []uuid.UUID) error
}

type MediaServicePCWithTransaction interface {
	MediaServicePC
	WithTransaction(db database.Database) MediaServicePC
}

type mediaServicePC struct {
	service *mediaService
}

func (pc *mediaServicePC) WithTransaction(db database.Database) MediaServicePC {
	return &mediaServicePC{
		service: &mediaService{
			storage:  pc.service.storage,
			database: db,
		},
	}
}

func NewMediaServicePC(params mediaServiceParams) MediaServicePCWithTransaction {
	log.Printf("creating %s-service PC", serviceName)
	return &mediaServicePC{
		service: newMediaService(params),
	}
}
