package media

import (
	"fmt"
	"log"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/services"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func getCompleteMultipartPath(mediaID uuid.UUID) string {
	return fmt.Sprintf("/media/%s/complete-multipart", mediaID.String())
}

type mediaServiceMux struct {
	service    *mediaService
	middleware core.MiddlewarePC
}

func (m *mediaServiceMux) BasePath() string {
	return "/media"
}

func (m *mediaServiceMux) Router() *chi.Mux {
	mux := chi.NewMux()

	mux.Group(func(router chi.Router) {
		router.Use(m.middleware.ValidateSignedURL)

		router.Post("/{mediaID}/post-processing", m.postProcessingMediaItems)
	})

	mux.Group(func(router chi.Router) {
		router.Use(m.middleware.ValidateAndGetActorDetailsFromHttpRequest)
		router.Use(m.middleware.ValidateActorTypeUserGlobalStatus)
		router.Use(m.middleware.ValidateSignedURLWithActor)

		router.Post("/{mediaID}/complete-multipart", m.service.completeMultipartUpload)
	})

	return mux
}

func NewMediaServiceMux(params mediaServiceMuxParams) services.Mux {
	log.Printf("adding %s-service mux", serviceName)
	return &mediaServiceMux{
		service:    newMediaService(params),
		middleware: params.Middleware(),
	}
}
