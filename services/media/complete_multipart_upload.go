package media

import (
	"context"
	"errors"
	"net/http"

	"github.com/Adgytec/adgytec-flow/config/storage"
	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type partsInfo struct {
	Etag       string `json:"etag"`
	PartNumber int32  `json:"partNumber"`
}

func (p partsInfo) GetEtag() string {
	return p.Etag
}

func (p partsInfo) GetPartNumber() int32 {
	return p.PartNumber
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

func (d completeMultipartUploadData) GetPartsInfo() []storage.MultipartPartInfo {
	parts := make([]storage.MultipartPartInfo, 0, len(d.PartsInfo))
	for _, part := range d.PartsInfo {
		parts = append(parts, part)
	}

	return parts
}

func (d completeMultipartUploadData) Validate() error {
	validationErr := validation.ValidateStruct(
		&d,
		validation.Field(
			// each slice item Validate() method is implicitly called
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

func (s *mediaService) completeMultipartUpload(ctx context.Context, mediaID uuid.UUID, multipartDetails completeMultipartUploadData) error {
	uploadDetails, dbErr := s.database.Queries().GetMediaItemDetails(ctx, mediaID)
	if dbErr != nil {
		if errors.Is(dbErr, pgx.ErrNoRows) {
			return &MediaItemNotFoundError{}
		}
		return dbErr

	}

	if uploadDetails.UploadType != db.GlobalMediaUploadTypeMultipart {
		return ErrUploadNotMultipart
	}

	if uploadDetails.UploadID == nil {
		return ErrMultipartUploadIDNotFound
	}

	completeMultipartErr := s.storage.CompleteMultipartUpload(ctx, uploadDetails.Key, *uploadDetails.UploadID, multipartDetails.GetPartsInfo())
	if completeMultipartErr != nil {
		return completeMultipartErr
	}

	return nil
}

func (s *mediaServiceMux) completeMultipartUpload(w http.ResponseWriter, r *http.Request) {
	reqCtx := r.Context()

	completeMultipartDetails, payloadErr := payload.DecodeRequestBodyAndValidate[completeMultipartUploadData](w, r)
	if payloadErr != nil {
		payload.EncodeError(w, payloadErr)
		return
	}

	mediaID := chi.URLParam(r, "mediaID")
	mediaUUID, mediaIDErr := s.service.getMediaUUIDFromString(mediaID)
	if mediaIDErr != nil {
		payload.EncodeError(w, mediaIDErr)
		return
	}

	completeMultipartErr := s.service.completeMultipartUpload(reqCtx, mediaUUID, completeMultipartDetails)
	if completeMultipartErr != nil {
		payload.EncodeError(w, completeMultipartErr)
		return
	}
}
