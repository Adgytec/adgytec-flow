package app_errors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
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

func (e *InvalidCursorValueError) HTTPResponse() core.ResponseHTTPError {
	return core.ResponseHTTPError{
		HTTPStatusCode: http.StatusBadRequest,
		Message:        helpers.ValuePtr(e.Error()),
	}
}
