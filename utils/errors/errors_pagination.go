package app_errors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/apires"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

var (
	ErrInvalidCursorValue = errors.New("invalid cursor value")
)

type InvalidCursorValueError struct {
	Cursor string
}

func (e *InvalidCursorValueError) Error() string {
	return fmt.Sprintf("Provided cursor value: '%s' is invalid.", e.Cursor)
}

func (e *InvalidCursorValueError) Is(target error) bool {
	return target == ErrInvalidCursorValue
}

func (e *InvalidCursorValueError) HTTPResponse() apires.ErrorDetails {
	return apires.ErrorDetails{
		HTTPStatusCode: http.StatusBadRequest,
		Message:        pointer.New(e.Error()),
	}
}
