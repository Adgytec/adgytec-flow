package media

import (
	"context"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/rs/zerolog/log"
)

func (s *mediaService) newMediaItems(ctx context.Context, input []NewMediaItemInfoWithStorageDetails) ([]MediaUploadDetails, error) {
	// check new media limit per action
	if len(input) < 1 || len(input) > int(mediaUploadLimit) {
		return nil, &InvalidNumberOfNewMediaItemsError{
			itemLength: len(input),
		}
	}

	// get upload details
	uploadDetails := make([]MediaUploadDetails, 0, len(input))
	mediaItemsParams := make([]db.AddMediaItemsParams, 0, len(input))
	for _, mediaItem := range input {
		mediaUploadDetail, mediaUploadDetailErr := s.getUploadDetails(ctx, mediaItem)
		if mediaUploadDetailErr != nil {
			return nil, mediaUploadDetailErr
		}

		newMediaItemParams := db.AddMediaItemsParams{
			ID:               mediaUploadDetail.ID,
			BucketPath:       mediaItem.getMediaItemKey(),
			RequiredMimeType: mediaItem.getRequiredMime(),
			UploadType:       mediaUploadDetail.UploadType,
			UploadID:         mediaUploadDetail.multipartUploadID,
		}

		uploadDetails = append(uploadDetails, *mediaUploadDetail)
		mediaItemsParams = append(mediaItemsParams, newMediaItemParams)
	}

	// add details to db
	_, dbErr := s.database.Queries().AddMediaItems(ctx, mediaItemsParams)
	if dbErr != nil {
		log.Error().
			Err(dbErr).
			Str("action", "adding new media items to db").
			Send()

		return nil, dbErr
	}

	return uploadDetails, nil
}

func (a *mediaServiceActions) NewMediaItem(ctx context.Context, input NewMediaItemInfoWithStorageDetails) (*MediaUploadDetails, error) {
	uploadDetails, newMediaErr := a.service.newMediaItems(ctx, []NewMediaItemInfoWithStorageDetails{input})
	if newMediaErr != nil {
		return nil, newMediaErr
	}

	if len(uploadDetails) != 1 {
		return nil, ErrCreatingNewMediaItem
	}

	return &uploadDetails[0], nil
}

func (a *mediaServiceActions) NewMediaItems(ctx context.Context, input []NewMediaItemInfoWithStorageDetails) ([]MediaUploadDetails, error) {
	return a.service.newMediaItems(ctx, input)
}
