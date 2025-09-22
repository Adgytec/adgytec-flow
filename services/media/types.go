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

type NewMediaItemInput struct {
	Size     int64
	Name     string
	MimeType string
}

// Validate() validates the input values
func (mediaItemInput NewMediaItemInput) Validate() error {
	validationErr := validation.ValidateStruct(&mediaItemInput,
		validation.Field(
			&mediaItemInput.Size,
			validation.Required,
			validation.Min(0),
		),
		validation.Field(
			&mediaItemInput.Name,
			validation.Required,
		),
		validation.Field(
			&mediaItemInput.MimeType,
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

func (mediaItemInput NewMediaItemInput) ensureMediaItemType(requiredType string) error {
	if !strings.HasPrefix(mediaItemInput.MimeType, requiredType) {
		return &InvalidMediaTypeValueError{
			Required: fmt.Sprintf("%s/*", requiredType),
			Got:      mediaItemInput.MimeType,
		}
	}

	return nil
}

// EnsureMediaItemIsImage() ensures the item that will be uploaded is image
func (mediaItemInput NewMediaItemInput) EnsureMediaItemIsImage() error {
	return mediaItemInput.ensureMediaItemType("image")
}

// EnsureMediaItemIsVideo() ensures the item that will be uploaded is video
func (mediaItemInput NewMediaItemInput) EnsureMediaItemIsVideo() error {
	return mediaItemInput.ensureMediaItemType("video")
}

func (mediaItemInput NewMediaItemInput) getMediaItemExtension() string {
	return filepath.Ext(mediaItemInput.Name)
}

type NewMediaItemInputWithBucketPrefix struct {
	NewMediaItemInput
	BucketPrefix string
}

func (mediaItemInput NewMediaItemInputWithBucketPrefix) getMediaItemKey() string {
	return path.Join(
		mediaItemInput.BucketPrefix,
		uuid.NewString()+mediaItemInput.getMediaItemExtension(),
	)
}

type MultipartPartUploadOutput struct {
	PresignPut string `json:"presignPut"`
	PartNumber int32  `json:"partNumber"`
	PartSize   int64  `json:"partSize"`
}

type NewMediaItemOutput struct {
	MediaID              uuid.UUID                   `json:"mediaID"`
	UploadType           db.GlobalMediaUploadType    `json:"uploadType"`
	PresignPut           *string                     `json:"presignPut,omitempty"`
	MultipartPresignPart []MultipartPartUploadOutput `json:"multipartPresignPart,omitempty"`
}
