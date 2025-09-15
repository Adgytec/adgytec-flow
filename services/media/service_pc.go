package media

import (
	"context"
	"log"

	"github.com/google/uuid"
)

type MediaServicePC interface {
	NewMediaItem(ctx context.Context, input NewMediaItemInput) (NewMediaItemOutput, error)
	NewMediaItems(ctx context.Context, input []NewMediaItemInput) ([]NewMediaItemOutput, error)
	CompleteMediaItemUpload(ctx context.Context, mediaID uuid.UUID) error
	CompleteMediaItemsUpload(ctx context.Context, mediaIDs []uuid.UUID) error
}

type mediaServicePC struct {
	service *mediaService
}

func NewMediaServicePC(params mediaServiceParams) MediaServicePC {
	log.Printf("creating %s-service PC", serviceName)
	return &mediaServicePC{
		service: newMediaService(params),
	}
}
