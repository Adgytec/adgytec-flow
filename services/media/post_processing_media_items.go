package media

import (
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/google/uuid"
)

func (s *mediaService) postProcessingMediaItems(mediaID uuid.UUID) error {
	return core.ErrNotImplemented
}

func (m *mediaServiceMux) postProcessingMediaItems(w http.ResponseWriter, r *http.Request) {

}
