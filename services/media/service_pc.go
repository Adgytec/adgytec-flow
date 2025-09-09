package media

import (
	"log"

	"github.com/google/uuid"
)

type MediaServicePC interface {
	UploadSuccess(mediaIDs []uuid.UUID) error
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
