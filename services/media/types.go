package media

import (
	"path"
	"path/filepath"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/core"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type NewMediaItem struct {
	Size         int64
	Name         string
	RequiredMime []string
}

// Validate() validates the input values
func (mediaItem NewMediaItem) Validate() error {
	validationErr := validation.ValidateStruct(&mediaItem,
		validation.Field(
			&mediaItem.Size,
			validation.Required,
			validation.Min(0),
		),
		validation.Field(
			&mediaItem.Name,
			validation.Required,
		),
		validation.Field(
			&mediaItem.RequiredMime,
			validation.Required,
			validation.Each(validation.Required),
		))

	if validationErr != nil {
		return &core.FieldValidationError{
			ValidationErrors: validationErr,
		}
	}

	return nil
}

func (mediaItem NewMediaItem) getMediaItemExtension() string {
	return filepath.Ext(mediaItem.Name)
}

type NewMediaItemWithBucketPrefix struct {
	NewMediaItem
	BucketPrefix string
}

func (mediaItem NewMediaItemWithBucketPrefix) getMediaItemKey() string {
	return path.Join(
		mediaItem.BucketPrefix,
		uuid.NewString()+mediaItem.getMediaItemExtension(),
	)
}

type MultipartPartUpload struct {
	PresignPut string `json:"presignPut"`
	PartNumber int32  `json:"partNumber"`
	PartSize   int64  `json:"partSize"`
}

type MediaUploadDetails struct {
	MediaID               uuid.UUID                  `json:"mediaID"`
	UploadType            db.GlobalMediaUploadType   `json:"uploadType"`
	PresignPut            *string                    `json:"presignPut,omitempty"`
	MultipartPresignPart  []MultipartPartUpload      `json:"multipartPresignPart,omitempty"`
	CompleteUploadActions MediaUploadCompleteActions `json:"completeUploadActions"`
}

type MediaUploadCompleteActions struct {
	Success string `json:"success"`
	Failed  string `json:"failed"`
}
