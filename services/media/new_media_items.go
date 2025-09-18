package media

import (
	"context"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

func (s *mediaService) newMediaItems(ctx context.Context, input []NewMediaItemInputWithBucketPrefix) ([]NewMediaItemOutput, error) {
	return nil, core.ErrNotImplemented
}

func (pc *mediaServicePC) NewMediaItems(ctx context.Context, input []NewMediaItemInputWithBucketPrefix) ([]NewMediaItemOutput, error) {
	return pc.service.newMediaItems(ctx, input)
}

func (pc *mediaServicePC) NewMediaItem(ctx context.Context, input NewMediaItemInputWithBucketPrefix) (*NewMediaItemOutput, error) {
	output, err := pc.service.newMediaItems(ctx, []NewMediaItemInputWithBucketPrefix{input})
	if err != nil {
		return nil, err
	}

	if len(output) != 1 {
		return nil, ErrCreatingNewMediaItem
	}

	return &output[0], nil
}
