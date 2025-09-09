package media

import (
	"log"

	"github.com/google/uuid"
)

type MediaServicePC interface {
	UploadSuccess(mediaIDs []uuid.UUID) error

	// TODO: this method parameter and return type will be updated when actual implementation are added
	NewPresignRequest(key string) error
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
