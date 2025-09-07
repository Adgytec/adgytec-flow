package pagination

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/apires"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

var (
	ErrInvalidCursorValue             = errors.New("invalid cursor value")
	ErrPaginationActionNotImplemented = errors.New("action not implemented")
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

type PaginationActionNotImplementedError struct {
	Action string
}

func (e *PaginationActionNotImplementedError) Error() string {
	if e.Action == "" {
		return ErrPaginationActionNotImplemented.Error()
	}

	return fmt.Sprintf("Action: %s, not implemented", e.Action)
}

func (e *PaginationActionNotImplementedError) Is(target error) bool {
	return target == ErrPaginationActionNotImplemented
}

func (e *PaginationActionNotImplementedError) HTTPResponse() apires.ErrorDetails {
	return apires.ErrorDetails{
		HTTPStatusCode: http.StatusNotImplemented,
		Message:        pointer.New(e.Error()),
	}
}
