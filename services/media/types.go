package media

import (
	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/core"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type NewMediaItemInput struct {
	Size      int64
	MediaType db.GlobalMediaType
	Name      string
}

// Validate() validates the input values
func (mediaItemInput NewMediaItemInput) Validate() error {
	validationErr := validation.ValidateStruct(&mediaItemInput,
		validation.Field(
			&mediaItemInput.Size,
			validation.Required,
			validation.Min(1),
		),
		validation.Field(
			&mediaItemInput.MediaType,
			validation.Required,
			validation.By(
				func(_ any) error {
					if !mediaItemInput.MediaType.Valid() {
						return ErrInvalidMediaTypeValue
					}
					return nil
				},
			),
		),
		validation.Field(
			&mediaItemInput.Name,
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

func (mediaItemInput NewMediaItemInput) ensureMediaTypeValue(value db.GlobalMediaType) error {
	validationErr := mediaItemInput.Validate()
	if validationErr != nil {
		return validationErr
	}

	mediaTypeValidationErr := validation.ValidateStruct(
		&mediaItemInput,
		validation.Field(
			&mediaItemInput.MediaType,
			validation.In(
				value,
			),
		),
	)

	if mediaTypeValidationErr != nil {
		return &core.FieldValidationError{
			ValidationErrors: mediaTypeValidationErr,
		}
	}

	return nil

}

// EnsureMediaItemIsImage() ensures the item that will be uploaded is image
func (mediaItemInput NewMediaItemInput) EnsureMediaItemIsImage() error {
	return mediaItemInput.ensureMediaTypeValue(db.GlobalMediaTypeImage)
}

// EnsureMediaItemIsVideo() ensures the item that will be uploaded is video
func (mediaItemInput NewMediaItemInput) EnsureMediaItemIsVideo() error {
	return mediaItemInput.ensureMediaTypeValue(db.GlobalMediaTypeVideo)
}

type NewMediaItemOutput struct{}
