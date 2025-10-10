package storage

import (
	"context"
	"log"
	"sort"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func (s *s3Client) CompleteMultipartUpload(ctx context.Context, key, uploadID string, partsInfo []MultipartPartInfo) error {
	// 1️⃣ Sort by PartNumber
	sort.Slice(partsInfo, func(i, j int) bool {
		return partsInfo[i].GetPartNumber() < partsInfo[j].GetPartNumber()
	})

	partDetails := make([]types.CompletedPart, 0, len(partsInfo))
	for i := 1; i <= len(partsInfo); i++ {
		part := partsInfo[i-1]
		if i != int(part.GetPartNumber()) {
			return &InvalidPartNumbersError{}
		}

		partDetails = append(partDetails, types.CompletedPart{
			ETag:       aws.String(part.GetEtag()),
			PartNumber: aws.Int32(part.GetPartNumber()),
		})
	}

	completeMultipart := types.CompletedMultipartUpload{
		Parts: partDetails,
	}

	_, completeUploadErr := s.client.CompleteMultipartUpload(
		ctx,
		&s3.CompleteMultipartUploadInput{
			Bucket:          aws.String(s.bucket),
			Key:             aws.String(key),
			UploadId:        aws.String(uploadID),
			MultipartUpload: &completeMultipart,
		},
	)
	if completeUploadErr != nil {
		log.Printf("error completing multipart upload for '%s': %v", key, completeUploadErr)
		return completeUploadErr
	}

	return nil
}
