package storage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func (s *s3Client) CompleteMultipartUpload(ctx context.Context, key, uploadID string, partsInfo types.CompletedMultipartUpload) error {
	_, completeUploadErr := s.client.CompleteMultipartUpload(
		ctx,
		&s3.CompleteMultipartUploadInput{
			Bucket:          aws.String(s.bucket),
			Key:             aws.String(key),
			UploadId:        aws.String(uploadID),
			MultipartUpload: &partsInfo,
		},
	)
	return completeUploadErr
}
