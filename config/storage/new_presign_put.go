package storage

import (
	"context"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

func (s *s3Client) NewPresignPut(ctx context.Context, key string) (string, error) {

	return "", core.ErrNotImplemented
}
