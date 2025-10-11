package storage

import (
	"context"

	"github.com/Adgytec/adgytec-flow/utils/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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
		logger.GetLoggerFromContext(ctx).Error().
			Err(presignErr).
			Int32("part-number", partNumber).
			Str("key", key).
			Str("action", "multipart upload part presign put generation").
			Send()

		return "", presignErr
	}

	return presignHTTPReq.URL, nil
}
