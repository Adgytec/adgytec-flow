package storage

import (
	"context"
	"log"

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
	if completeUploadErr != nil {
		log.Printf("error completing multipart upload for '%s': %v", key, completeUploadErr)
		return completeUploadErr
	}

	return nil
}
