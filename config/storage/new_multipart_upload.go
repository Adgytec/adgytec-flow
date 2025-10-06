package storage

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
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
		log.Printf("error creating multipart upload for '%s': %v", key, newMultipartUploadErr)
		return "", newMultipartUploadErr
	}

	return *newMultipartUploadOutput.UploadId, nil
}
