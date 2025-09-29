package storage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *s3Client) NewPresignPut(ctx context.Context, key string) (string, error) {
	presignHTTPReq, presignErr := s.presignClient.PresignPutObject(
		ctx,
		&s3.PutObjectInput{
			Bucket: aws.String(s.bucket),
			Key:    aws.String(key),
		},
		func(po *s3.PresignOptions) {
			po.Expires = s.presignExpiration
		},
	)
	if presignErr != nil {
		return "", presignErr
	}

	return presignHTTPReq.URL, nil
}
