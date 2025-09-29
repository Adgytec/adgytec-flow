package storage

import (
	"context"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

func (s *s3Client) NewPresignUploadPart(ctx context.Context, key, uploadID string, partNumber int32) (string, error) {
	return "", core.ErrNotImplemented
}
