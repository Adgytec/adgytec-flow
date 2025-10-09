package media

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/types"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type NewMediaItemInfo struct {
	ID   uuid.UUID `json:"id"`
	Size uint64    `json:"size"`
	Name string    `json:"name"`
}

// NullableNewMediaItemInfo is used with request bodies
type NullableNewMediaItemInfo = types.Nullable[NewMediaItemInfo]

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
			&mediaItem.Name,
			validation.Required,
		),
	)

	if validationErr != nil {
		return validationErr
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

type MediaUploadDetails struct {
	ID                       uuid.UUID                `json:"mediaID"`
	UploadType               db.GlobalMediaUploadType `json:"uploadType"`
	PresignPut               *string                  `json:"presignPut,omitempty"`
	MultipartPresignPart     []MultipartPartUpload    `json:"multipartPresignPart,omitempty"`
	MultipartSuccessCallback *string                  `json:"multipartSuccessCallback,omitempty"`
	multipartUploadID        *string
}

type MultipartPartUpload struct {
	PresignPut string `json:"presignPut"`
	PartNumber uint16 `json:"partNumber"`
	PartSize   uint64 `json:"partSize"`
}
