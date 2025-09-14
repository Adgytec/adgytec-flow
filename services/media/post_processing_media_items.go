package media

import (
	"net/http"

	"github.com/google/uuid"
)

func (s *mediaService) postProcessingMediaItems(mediaID uuid.UUID) error {
	return nil
}

func (m *mediaServiceMux) postProcessingMediaItems(w http.ResponseWriter, r *http.Request) {

}
