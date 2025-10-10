package media

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/apires"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

var (
	ErrInvalidMediaTypeValue        = errors.New("invalid media type value")
	ErrCreatingNewMediaItem         = errors.New("error creating new media item")
	ErrMediaIDGeneration            = errors.New("error generating media id")
	ErrInvalidNumberOfNewMediaItems = errors.New("invalid number of new media item")
	ErrInvalidMediaSize             = errors.New("invalid media size")
	ErrMediaTooLarge                = errors.New("media too large")
	ErrInvalidMediaID               = errors.New("invalid media id")
	ErrUploadNotMultipart           = errors.New("upload not multipart")
	ErrMediaItemNotFound            = errors.New("media item not found")
	ErrMultipartUploadIDNotFound    = errors.New("multipart upload id not found")
)

type MediaTooLargeError struct {
	Size uint64
}

func (e *MediaTooLargeError) Error() string {
	return fmt.Sprintf("Max upload limit of %d bytes, but got size of %d bytes", multipartUploadLimit, e.Size)
}

func (e *MediaTooLargeError) Is(target error) bool {
	return target == ErrMediaTooLarge
}

func (e *MediaTooLargeError) HTTPResponse() apires.ErrorDetails {
	return apires.ErrorDetails{
		HTTPStatusCode: http.StatusBadRequest,
		Message:        pointer.New(e.Error()),
	}
}

type InvalidMediaTypeValueError struct {
	Required string
	Got      string
}

func (e *InvalidMediaTypeValueError) Error() string {
	return fmt.Sprintf("Required media type: %s, got: %s", e.Required, e.Got)
}

func (e *InvalidMediaTypeValueError) Is(target error) bool {
	return target == ErrInvalidMediaTypeValue
}

func (e *InvalidMediaTypeValueError) HTTPResponse() apires.ErrorDetails {
	return apires.ErrorDetails{
		HTTPStatusCode: http.StatusBadRequest,
		Message:        pointer.New(e.Error()),
	}
}

type InvalidNumberOfNewMediaItemsError struct {
	itemLength int
}

func (e *InvalidNumberOfNewMediaItemsError) Error() string {
	return fmt.Sprintf("supported new media item per action in range of %d to %d, but got %d", 1, mediaUploadLimit, e.itemLength)
}

func (e *InvalidNumberOfNewMediaItemsError) Is(target error) bool {
	return target == ErrInvalidNumberOfNewMediaItems
}

func (e *InvalidNumberOfNewMediaItemsError) HTTPResponse() apires.ErrorDetails {
	return apires.ErrorDetails{
		HTTPStatusCode: http.StatusBadRequest,
		Message:        pointer.New(e.Error()),
	}
}

type MediaItemNotFoundError struct{}

func (e *MediaItemNotFoundError) Error() string {
	return ErrMediaItemNotFound.Error()
}

func (e *MediaItemNotFoundError) Is(target error) bool {
	return target == ErrMediaItemNotFound
}

func (e *MediaItemNotFoundError) HTTPResponse() apires.ErrorDetails {
	return apires.ErrorDetails{
		HTTPStatusCode: http.StatusNotFound,
		Message:        pointer.New(e.Error()),
	}
}
