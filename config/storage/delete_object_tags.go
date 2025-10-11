package storage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/zerolog/log"
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
		log.Error().
			Err(deleteTagErr).
			Str("key", key).
			Str("action", "delete s3 object tags").
			Send()
		return deleteTagErr
	}
	return nil
}
