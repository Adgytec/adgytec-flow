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
)

type MediaTooLargeError struct {
	Size int64
}

func (e *MediaTooLargeError) Error() string {
	return fmt.Sprintf("Max upload limit of %d bytes, but got size of % bytes", multipartUploadLimit, e.Size)
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
