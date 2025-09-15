package media

import (
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/google/uuid"
)

func (s *mediaService) completeMediaItemsUpload(mediaIDs []uuid.UUID) error {
	return core.ErrNotImplemented
}

func (pc *mediaServicePC) CompleteMediaItemsUpload(mediaIDs []uuid.UUID) error {
	return pc.service.completeMediaItemsUpload(mediaIDs)
}

func (pc *mediaServicePC) CompleteMediaItemUpload(mediaIDs uuid.UUID) error {
	return core.ErrNotImplemented
}
