package media

import (
	"log"

	"github.com/google/uuid"
)

type MediaServicePC interface {
	NewMediaItem(mediaIDs []uuid.UUID) error
	CompleteMediaUpload(key string) error
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
