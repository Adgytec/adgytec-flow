package media

import (
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/core"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type partsInfo struct {
	Etag       string `json:"etag"`
	PartNumber int32  `json:"partNumber"`
}

func (p partsInfo) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Etag, validation.Required),
		validation.Field(&p.PartNumber, validation.Required, validation.Min(1), validation.Max(int32(maximumPartsCount))),
	)
}

type completeMultipartUploadData struct {
	PartsInfo []partsInfo `json:"partsInfo"`
}

func (d *completeMultipartUploadData) Validate() error {
	validationErr := validation.ValidateStruct(
		d,
		validation.Field(
			&d.PartsInfo,
			validation.Required,
			validation.Length(int(minimumPartsCount), int(maximumPartsCount)),
		),
	)

	if validationErr != nil {
		return &core.FieldValidationError{
			ValidationErrors: validationErr,
		}
	}

	return nil
}

func (s *mediaService) completeMultipartUpload(w http.ResponseWriter, r *http.Request) {}
