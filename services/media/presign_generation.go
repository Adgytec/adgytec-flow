package media

import (
	"context"
	"time"

	"github.com/Adgytec/adgytec-flow/database/db"
)

const (
	completeMultipartPresignValidDuration = 4 * time.Hour
)

func (s *mediaService) getUploadDetails(ctx context.Context, newMediaItem NewMediaItemInfoWithStorageDetails) (*MediaUploadDetails, error) {
	itemUploadInfo, itemUploadInfoErr := getUploadInfo(newMediaItem.Size)
	if itemUploadInfoErr != nil {
		return nil, itemUploadInfoErr
	}

	uploadDetails := MediaUploadDetails{
		ID:         newMediaItem.ID,
		UploadType: itemUploadInfo.uploadType,
	}

	// handle singlepart upload
	if itemUploadInfo.uploadType == db.GlobalMediaUploadTypeSinglepart {
		presignPut, presignErr := s.storage.NewPresignPut(ctx, newMediaItem.getMediaItemKey(), newMediaItem.ID)
		if presignErr != nil {
			return nil, presignErr
		}

		uploadDetails.PresignPut = &presignPut
	} else {
		// handle multipart upload
		multipartUploadID, multipartErr := s.storage.NewMultipartUpload(ctx, newMediaItem.getMediaItemKey(), newMediaItem.ID)
		if multipartErr != nil {
			return nil, multipartErr
		}

		// generate presign for each part
		partUploadDetails := make([]MultipartPartUpload, 0, itemUploadInfo.partCount)
		for i := 1; i <= int(itemUploadInfo.partCount); i++ {
			multipartPresignPart, presignErr := s.storage.NewPresignUploadPart(ctx, newMediaItem.getMediaItemKey(), multipartUploadID, int32(i))
			if presignErr != nil {
				return nil, presignErr
			}

			// last part size handling
			partSize := itemUploadInfo.partSize
			if i == int(itemUploadInfo.partCount) {
				partSize = itemUploadInfo.lastPartSize
			}

			part := MultipartPartUpload{
				PartNumber: uint16(i),
				PartSize:   partSize,
				PresignPut: multipartPresignPart,
			}
			partUploadDetails = append(partUploadDetails, part)
		}

		completeMultipartSignedURL, presignErr := s.newCompleteMultipartSignedURL(ctx, multipartUploadID)
		if presignErr != nil {
			return nil, presignErr
		}

		uploadDetails.MultipartPresignPart = partUploadDetails
		uploadDetails.MultipartSuccessCallback = &completeMultipartSignedURL
	}

	return &uploadDetails, nil
}

func (s *mediaService) newCompleteMultipartSignedURL(ctx context.Context, uploadID string) (string, error) {
	signedURL, err := s.auth.NewSignedURLWithActor(ctx, getCompleteMultipartPath(uploadID), nil, completeMultipartPresignValidDuration)
	if err != nil {
		return "", err
	}

	return signedURL.String(), nil
}
