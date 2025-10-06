package media

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/apires"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

var (
	ErrInvalidMediaTypeValue       = errors.New("invalid media type value")
	ErrCreatingNewMediaItem        = errors.New("error creating new media item")
	ErrMediaIDGeneration           = errors.New("error generating media id")
	ErrInvalidNumberOfNewMediaItem = errors.New("invalid number of new media item")
	ErrInvalidMediaSize            = errors.New("invalid media size")
	ErrMediaTooLarge               = errors.New("media too large")
	ErrMultipartTooSmall           = errors.New("media too small for multipart upload")
)

type MediaTooLargeError struct {
	Size int
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
