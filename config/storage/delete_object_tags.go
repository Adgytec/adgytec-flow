package storage

import (
	"context"
	"log"

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
		log.Printf("error deleting s3 object tags for '%s': %v", key, deleteTagErr)
		return deleteTagErr
	}
	return nil
}
