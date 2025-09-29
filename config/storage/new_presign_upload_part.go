package storage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *s3Client) NewPresignUploadPart(ctx context.Context, key, uploadID string, partNumber int32) (string, error) {
	presignHTTPReq, presingErr := s.presignClient.PresignUploadPart(
		ctx,
		&s3.UploadPartInput{
			Bucket:     aws.String(s.bucket),
			Key:        aws.String(key),
			PartNumber: aws.Int32(partNumber),
			UploadId:   aws.String(uploadID),
		},
		func(po *s3.PresignOptions) {
			po.Expires = s.presignExpiration
		},
	)
	if presingErr != nil {
		return "", presingErr
	}

	return presignHTTPReq.URL, nil
}
