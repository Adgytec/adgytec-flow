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
	WithTransaction(db database.Database) MediaServicePC
}

type mediaServicePC struct {
	service *mediaService
}

func (pc *mediaServicePC) WithTransaction(db database.Database) MediaServicePC {
	mediaServiceCopy := *pc.service
	mediaServiceCopy.database = db

	return &mediaServicePC{
		service: &mediaServiceCopy,
	}
}

func NewMediaServicePC(params mediaServiceParams) MediaServicePC {
	log.Printf("creating %s-service PC", serviceName)
	return &mediaServicePC{
		service: newMediaService(params),
	}
}
