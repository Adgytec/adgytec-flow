package storage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/zerolog/log"
)

func (s *s3Client) NewPresignUploadPart(ctx context.Context, key, uploadID string, partNumber int32) (string, error) {
	presignHTTPReq, presignErr := s.presignClient.PresignUploadPart(
		ctx,
		&s3.UploadPartInput{
			Bucket:     aws.String(s.bucket),
			Key:        aws.String(key),
			PartNumber: aws.Int32(partNumber),
			UploadId:   aws.String(uploadID),
		},
		func(po *s3.PresignOptions) {
			po.Expires = presignExpiration
		},
	)
	if presignErr != nil {
		log.Error().
			Err(presignErr).
			Int32("part-number", partNumber).
			Str("key", key).
			Str("action", "multipart upload part presign put generation").
			Send()

		return "", presignErr
	}

	return presignHTTPReq.URL, nil
}
