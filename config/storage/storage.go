package storage

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/google/uuid"
)

const presignExpiration = time.Hour

type Storage interface {
	NewPresignPut(ctx context.Context, key string, id uuid.UUID) (string, error)
	NewMultipartUpload(ctx context.Context, key string, id uuid.UUID) (string, error)
	NewPresignUploadPart(ctx context.Context, key, uploadID string, partNumber int32) (string, error)
	CompleteMultipartUpload(ctx context.Context, key, uploadID string, partsInfo types.CompletedMultipartUpload) error
	DeleteObjectTags(ctx context.Context, key string) error
}
