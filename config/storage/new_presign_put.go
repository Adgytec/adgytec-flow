package storage

import (
	"context"

	"github.com/Adgytec/adgytec-flow/utils/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

func (s *s3Client) NewPresignPut(ctx context.Context, key string, id uuid.UUID) (string, error) {
	presignHTTPReq, presignErr := s.presignClient.PresignPutObject(
		ctx,
		&s3.PutObjectInput{
			Bucket:  aws.String(s.bucket),
			Key:     aws.String(key),
			Tagging: aws.String(newObjectTag(id)),
		},
		func(po *s3.PresignOptions) {
			po.Expires = presignExpiration
		},
	)
	if presignErr != nil {
		logger.GetLoggerFromContext(ctx).Error().
			Err(presignErr).
			Str("key", key).
			Str("action", "new presign put").
			Send()
		return "", presignErr
	}

	return presignHTTPReq.URL, nil
}
