package media

import (
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/google/uuid"
)

func (s *mediaService) uploadSuccess(mediaIDs []uuid.UUID) error {
	return core.ErrNotImplemented
}

func (pc *mediaServicePC) UploadSuccess(mediaIDs []uuid.UUID) error {
	return pc.service.uploadSuccess(mediaIDs)
}
