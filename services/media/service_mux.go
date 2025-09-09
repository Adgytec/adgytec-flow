package media

import (
	"log"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/services"
	"github.com/go-chi/chi/v5"
)

type mediaServiceMux struct {
	service    *mediaService
	middleware core.MiddlewarePC
}

func (m *mediaServiceMux) BasePath() string {
	return "/media"
}

func (m *mediaServiceMux) Router() *chi.Mux {
	mux := chi.NewMux()

	// add webhooks methods for media pipeline

	return mux
}

func NewMediaServiceMux(params mediaServiceMuxParams) services.Mux {
	log.Printf("adding %s-service mux", serviceName)
	return &mediaServiceMux{
		service:    newMediaService(params),
		middleware: params.Middleware(),
	}
}
