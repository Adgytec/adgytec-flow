package media

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/core"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type NewMediaItemInfo struct {
	ID   uuid.UUID
	Size int64
	Name string
}

// Validate() validates the input values
func (mediaItem NewMediaItemInfo) Validate() error {
	validationErr := validation.ValidateStruct(&mediaItem,
		validation.Field(
			&mediaItem.ID,
			validation.Required,
			validation.By(func(val any) error {
				id := val.(uuid.UUID)
				if id == uuid.Nil {
					return fmt.Errorf("id cannot be nil UUID")
				}
				return nil
			}),
		),
		validation.Field(
			&mediaItem.Size,
			validation.Required,
			validation.Min(0),
		),
		validation.Field(
			&mediaItem.Name,
			validation.Required,
		),
	)

	if validationErr != nil {
		return &core.FieldValidationError{
			ValidationErrors: validationErr,
		}
	}

	return nil
}

type NewMediaItemInfoWithStorageDetails struct {
	NewMediaItemInfo
	BucketPrefix string
	RequiredMime []string
}

func (mediaItem NewMediaItemInfoWithStorageDetails) getMediaItemKey() string {
	return path.Join(
		mediaItem.BucketPrefix,
		mediaItem.ID.String()+filepath.Ext(mediaItem.Name),
	)
}

func (mediaItem NewMediaItemInfoWithStorageDetails) getRequiredMime() []string {
	zero := []string{zeroMime}
	requiredMime := make([]string, 0, len(mediaItem.RequiredMime))

	for _, m := range mediaItem.RequiredMime {
		m = strings.TrimSpace(m)
		if m == "" {
			continue // skip empty
		}

		requiredMime = append(requiredMime, m)
	}

	if len(requiredMime) == 0 {
		return zero
	}

	return requiredMime
}

type MultipartPartUpload struct {
	PresignPut string `json:"presignPut"`
	PartNumber int32  `json:"partNumber"`
	PartSize   int64  `json:"partSize"`
}

type MediaUploadDetails struct {
	ID                    uuid.UUID                  `json:"mediaID"`
	UploadType            db.GlobalMediaUploadType   `json:"uploadType"`
	PresignPut            *string                    `json:"presignPut,omitempty"`
	MultipartPresignPart  []MultipartPartUpload      `json:"multipartPresignPart,omitempty"`
	CompleteUploadActions MediaUploadCompleteActions `json:"completeUploadActions"`
}

type MediaUploadCompleteActions struct {
	Success string `json:"success"`
	Failed  string `json:"failed"`
}
