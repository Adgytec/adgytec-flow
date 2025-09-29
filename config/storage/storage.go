package storage

import "context"

type Storage interface {
	NewPresignPut(ctx context.Context, key string) (string, error)
	NewMultipartUpload(ctx context.Context, key string) (string, error)
	NewPresignUploadPart(ctx context.Context, key, uploadID string, partNumber int32) (string, error)
	CompleteMultipartUpload(ctx context.Context, key, uploadID string) error
}
