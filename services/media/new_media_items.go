package media

import (
	"context"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
	"github.com/google/uuid"
)

func (s *mediaService) validateNewMediaItemCount(input []NewMediaItemInputWithBucketPrefix) error {
	if len(input) < 1 || len(input) > mediaUploadLimit {
		return ErrInvalidNumberOfNewMediaItem
	}
	return nil
}

func (s *mediaService) prepareMediaItems(
	input []NewMediaItemInputWithBucketPrefix,
) ([]NewMediaItemOutput, []db.NewTemporaryMediaParams, error) {

	outputs := make([]NewMediaItemOutput, 0, len(input))
	dbParams := make([]db.NewTemporaryMediaParams, 0, len(input))

	for _, val := range input {
		// check size
		if val.Size > multipartUploadLimit {
			return nil, nil, &MediaTooLargeError{
				Size: val.Size,
			}
		}

		// generate IDs
		mediaID, err := uuid.NewV7()
		if err != nil {
			return nil, nil, ErrMediaIDGeneration
		}
		mediaKey := val.getMediaItemKey()

		// decide upload type
		output, param, err := s.prepareSingleMediaItem(mediaID, mediaKey, val)
		if err != nil {
			return nil, nil, err
		}

		outputs = append(outputs, output)
		dbParams = append(dbParams, param)
	}

	return outputs, dbParams, nil
}

func (s *mediaService) prepareSingleMediaItem(
	mediaID uuid.UUID,
	mediaKey string,
	val NewMediaItemInputWithBucketPrefix,
) (NewMediaItemOutput, db.NewTemporaryMediaParams, error) {
	output := NewMediaItemOutput{
		MediaID: mediaID,
	}
	var uploadID *string

	if val.Size >= singlepartUploadLimit {
		// multipart upload
		output.UploadType = db.GlobalMediaUploadTypeMultipart
		upload, err := s.prepareMultipartUpload(mediaKey, val.Size)
		if err != nil {
			return output, db.NewTemporaryMediaParams{}, err
		}

		output.MultipartPresignPart = upload.parts
		uploadID = &upload.id

	} else {
		// singlepart upload
		output.UploadType = db.GlobalMediaUploadTypeSinglepart
		presignURL, err := s.storage.NewPresignPut(mediaKey)
		if err != nil {
			return output, db.NewTemporaryMediaParams{}, err
		}

		output.PresignPut = pointer.New(presignURL)
	}

	param := db.NewTemporaryMediaParams{
		ID:         mediaID,
		BucketPath: mediaKey,
		UploadType: output.UploadType,
		MediaType:  val.MediaType,
		UploadID:   uploadID,
	}

	return output, param, nil
}

type multipartUploadResult struct {
	id    string
	parts []MultipartPartUploadOutput
}

func (s *mediaService) prepareMultipartUpload(
	mediaKey string,
	size int64,
) (multipartUploadResult, error) {
	uploadID, err := s.storage.NewMultipartUpload(mediaKey)
	if err != nil {
		return multipartUploadResult{}, err
	}

	valSize := size
	partsCount := (size + multipartPartSize - 1) / multipartPartSize
	parts := make([]MultipartPartUploadOutput, 0, partsCount)

	for part := 1; part <= int(partsCount); part++ {
		partSize := min(multipartPartSize, valSize)
		valSize -= partSize

		presignURL, err := s.storage.NewPresignUploadPart(mediaKey, uploadID, int32(part))
		if err != nil {
			return multipartUploadResult{}, err
		}

		parts = append(parts, MultipartPartUploadOutput{
			PartNumber: int32(part),
			PartSize:   partSize,
			PresignPut: presignURL,
		})
	}

	return multipartUploadResult{id: uploadID, parts: parts}, nil
}

func (s *mediaService) saveTemporaryMedia(
	ctx context.Context,
	params []db.NewTemporaryMediaParams,
) error {
	qtx, tx, err := s.database.WithTransaction(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	if _, err := qtx.Queries().NewTemporaryMedia(ctx, params); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// unfinished uploads and multipart upload are manged using s3 lifecycle methods
func (s *mediaService) newMediaItems(ctx context.Context, input []NewMediaItemInputWithBucketPrefix) ([]NewMediaItemOutput, error) {
	itemCountErr := s.validateNewMediaItemCount(input)
	if itemCountErr != nil {
		return nil, itemCountErr
	}

	outputs, dbParams, err := s.prepareMediaItems(input)
	if err != nil {
		return nil, err
	}

	if err := s.saveTemporaryMedia(ctx, dbParams); err != nil {
		return nil, err
	}

	return outputs, nil
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
