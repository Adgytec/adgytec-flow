package media

import "context"

func (s *mediaService) newMediaItems(ctx context.Context, input []NewMediaItemInfoWithStorageDetails) ([]MediaUploadDetails, error) {
	return nil, nil
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
