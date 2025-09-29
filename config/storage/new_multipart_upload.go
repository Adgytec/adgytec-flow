package storage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *s3Client) NewMultipartUpload(ctx context.Context, key string) (string, error) {
	newMultipartUploadOutput, newMultipartUploadErr := s.client.CreateMultipartUpload(
		ctx,
		&s3.CreateMultipartUploadInput{
			Bucket:  aws.String(s.bucket),
			Key:     aws.String(key),
			Tagging: aws.String("status=temp"),
		},
	)
	if newMultipartUploadErr != nil {
		return "", newMultipartUploadErr
	}

	return *newMultipartUploadOutput.UploadId, nil
}
