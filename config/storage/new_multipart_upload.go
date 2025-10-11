package storage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (s *s3Client) NewMultipartUpload(ctx context.Context, key string, id uuid.UUID) (string, error) {
	newMultipartUploadOutput, newMultipartUploadErr := s.client.CreateMultipartUpload(
		ctx,
		&s3.CreateMultipartUploadInput{
			Bucket:  aws.String(s.bucket),
			Key:     aws.String(key),
			Tagging: aws.String(newObjectTag(id)),
		},
	)
	if newMultipartUploadErr != nil {
		log.Error().
			Err(newMultipartUploadErr).
			Str("key", key).
			Str("action", "new multipart upload").
			Send()
		return "", newMultipartUploadErr
	}

	return *newMultipartUploadOutput.UploadId, nil
}
