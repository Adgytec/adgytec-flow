package media

import (
	"context"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

func (s *mediaService) newMediaItems(ctx context.Context, input []NewMediaItemInput) ([]NewMediaItemOutput, error) {
	return nil, core.ErrNotImplemented
}

func (pc *mediaServicePC) NewMediaItems(ctx context.Context, input []NewMediaItemInput) ([]NewMediaItemOutput, error) {
	return pc.service.newMediaItems(ctx, input)
}

func (pc *mediaServicePC) NewMediaItem(ctx context.Context, input NewMediaItemInput) (NewMediaItemOutput, error) {
	var zero NewMediaItemOutput
	return zero, core.ErrNotImplemented
}
