package media

import "log"

type MediaServicePC interface{}

type mediaServicePC struct {
	service *mediaService
}

func NewMediaServicePC(params mediaServiceParams) MediaServicePC {
	log.Println("creating media-service PC")
	return &mediaServicePC{
		service: newMediaService(params),
	}
}
