package storage

import (
	"context"

	"github.com/Adgytec/adgytec-flow/utils/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *s3Client) DeleteObjectTags(ctx context.Context, key string) error {
	_, deleteTagErr := s.client.DeleteObjectTagging(
		ctx,
		&s3.DeleteObjectTaggingInput{
			Bucket: aws.String(s.bucket),
			Key:    aws.String(key),
		},
	)
	if deleteTagErr != nil {
		logger.GetLoggerFromContext(ctx).Error().
			Err(deleteTagErr).
			Str("key", key).
			Str("action", "delete s3 object tags").
			Send()
		return deleteTagErr
	}
	return nil
}
