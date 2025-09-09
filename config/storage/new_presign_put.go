package storage

import "github.com/Adgytec/adgytec-flow/utils/core"

func (s *s3Client) NewPresignPut(key string) (string, error) {
	return "", core.ErrNotImplemented
}
