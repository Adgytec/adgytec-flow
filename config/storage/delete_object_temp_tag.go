package storage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *s3Client) DeleteObjectTempTag(ctx context.Context, key string) error {
	_, deleteTagErr := s.client.DeleteObjectTagging(
		ctx,
		&s3.DeleteObjectTaggingInput{
			Bucket: aws.String(s.bucket),
			Key:    aws.String(key),
		},
	)
	return deleteTagErr
}
