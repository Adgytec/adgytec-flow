package storage

import "github.com/Adgytec/adgytec-flow/utils/core"

func (s *s3Client) AbortMultipartUpload(key, uploadID string) error {
	return core.ErrNotImplemented
}
