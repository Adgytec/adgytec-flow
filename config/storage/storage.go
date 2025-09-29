package storage

import (
	"context"
	"time"
)

const tempObjectTag = "status=temp"
const presignExpiration = time.Hour

type Storage interface {
	NewPresignPut(ctx context.Context, key string) (string, error)
	NewMultipartUpload(ctx context.Context, key string) (string, error)
	NewPresignUploadPart(ctx context.Context, key, uploadID string, partNumber int32) (string, error)
	CompleteMultipartUpload(ctx context.Context, key, uploadID string) error
	DeleteObjectTempTag(ctx context.Context, key string) error
}
