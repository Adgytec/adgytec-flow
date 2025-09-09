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
	log.Println("creating media-service PC")
	return &mediaServicePC{
		service: newMediaService(params),
	}
}
