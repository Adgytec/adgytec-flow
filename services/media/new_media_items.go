package media

import (
	"context"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
	"github.com/google/uuid"
)

func (s *mediaService) newMediaItems(ctx context.Context, input []NewMediaItemInputWithBucketPrefix) ([]NewMediaItemOutput, error) {
	if len(input) < 1 || len(input) > mediaUploadLimit {
		return nil, ErrInvalidNumberOfNewMediaItem
	}

	var newMediaItemOuput []NewMediaItemOutput
	var tempMediaParams []db.NewTemporaryMediaParams

	for _, val := range input {

		if val.Size > multipartUploadLimit {
			return nil, &MediaTooLargeError{
				Size: val.Size,
			}
		}

		mediaItemKey := val.getMediaItemKey()
		var uploadID *string

		mediaID, idErr := uuid.NewV7()
		if idErr != nil {
			return nil, ErrMediaIDGeneration
		}
		itemOutput := NewMediaItemOutput{
			MediaID: mediaID,
		}

		if val.Size >= singlepartUploadLimit {
			itemOutput.UploadType = db.GlobalMediaUploadTypeMultipart
			multipartUploadID, multipartErr := s.storage.NewMultipartUpload(mediaItemKey)
			if multipartErr != nil {
				return nil, multipartErr
			}

			uploadID = &multipartUploadID
			var uploadParts []MultipartPartUploadOutput
			partsCount := (val.Size + multipartPartSize - 1) / multipartPartSize
			valSize := val.Size

			for part := 1; part <= int(partsCount); part++ {
				if valSize < 1 {
					return nil, ErrInvalidMediaSize
				}

				partSize := multipartPartSize
				if valSize < multipartPartSize {
					partSize = int(valSize)
				}
				valSize -= int64(partSize)
				partDetail := MultipartPartUploadOutput{
					PartNumber: int32(part),
					PartSize:   int64(partSize),
				}

				presignURL, presignErr := s.storage.NewPresignUploadPart(mediaItemKey, multipartUploadID, int32(part))
				if presignErr != nil {
					return nil, presignErr
				}

				partDetail.PresignPut = presignURL

				uploadParts = append(uploadParts, partDetail)
			}

			itemOutput.MultipartPresignPart = uploadParts
		} else {
			itemOutput.UploadType = db.GlobalMediaUploadTypeSinglepart
			presignURL, presignErr := s.storage.NewPresignPut(mediaItemKey)
			if presignErr != nil {
				return nil, presignErr
			}

			itemOutput.PresignPut = pointer.New(presignURL)
		}

		tempMediaParams = append(tempMediaParams, db.NewTemporaryMediaParams{
			ID:         mediaID,
			BucketPath: mediaItemKey,
			UploadType: itemOutput.UploadType,
			MediaType:  val.MediaType,
			UploadID:   uploadID,
		})
		newMediaItemOuput = append(newMediaItemOuput, itemOutput)
	}

	qtx, tx, err := s.database.WithTransaction(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	_, dbErr := qtx.Queries().NewTemporaryMedia(ctx, tempMediaParams)
	if dbErr != nil {
		return nil, dbErr
	}

	commitErr := tx.Commit(ctx)
	if commitErr != nil {
		return nil, commitErr
	}
	return newMediaItemOuput, nil
}

func (pc *mediaServicePC) NewMediaItems(ctx context.Context, input []NewMediaItemInputWithBucketPrefix) ([]NewMediaItemOutput, error) {
	return pc.service.newMediaItems(ctx, input)
}

func (pc *mediaServicePC) NewMediaItem(ctx context.Context, input NewMediaItemInputWithBucketPrefix) (*NewMediaItemOutput, error) {
	output, err := pc.service.newMediaItems(ctx, []NewMediaItemInputWithBucketPrefix{input})
	if err != nil {
		return nil, err
	}

	if len(output) != 1 {
		return nil, ErrCreatingNewMediaItem
	}

	return &output[0], nil
}
