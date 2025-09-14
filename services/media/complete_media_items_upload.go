package media

import "github.com/google/uuid"

func (s *mediaService) completeMediaItemsUpload(mediaIDs []uuid.UUID) error {
	return nil
}

func (pc *mediaServicePC) CompleteMediaItemsUpload(mediaIDs []uuid.UUID) error {
	return pc.service.completeMediaItemsUpload(mediaIDs)
}
