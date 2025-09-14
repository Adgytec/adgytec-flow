package storage

import "github.com/Adgytec/adgytec-flow/utils/core"

func (s *s3Client) CompleteMultipartUpload(key, uploadID string) error {
	return core.ErrNotImplemented
}
