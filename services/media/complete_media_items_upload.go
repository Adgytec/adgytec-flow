package media

import (
	"context"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/google/uuid"
)

func (s *mediaService) completeMediaItemsUpload(ctx context.Context, mediaIDs []uuid.UUID) error {
	return core.ErrNotImplemented
}

func (pc *mediaServicePC) CompleteMediaItemsUpload(ctx context.Context, mediaIDs []uuid.UUID) error {
	return pc.service.completeMediaItemsUpload(ctx, mediaIDs)
}

func (pc *mediaServicePC) CompleteMediaItemUpload(ctx context.Context, mediaIDs uuid.UUID) error {
	return core.ErrNotImplemented
}
