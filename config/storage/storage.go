package storage

import "context"

type Storage interface {
	NewPresignPut(ctx context.Context, key string) (string, error)
	NewMultipartUpload(key string) (string, error)
	NewPresignUploadPart(key, uploadID string, partNumber int32) (string, error)
	CompleteMultipartUpload(key, uploadID string) error
	AbortMultipartUpload(key, uploadID string) error
	DeleteObject(key string) error
}
