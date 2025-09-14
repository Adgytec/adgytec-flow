package storage

import "github.com/Adgytec/adgytec-flow/utils/core"

func (s *s3Client) NewPresignUploadPart(key, uploadID string, partNumber int32) (string, error) {
	return "", core.ErrNotImplemented
}
