package core

import (
	"errors"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/apires"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

var (
	ErrNotImplemented           = errors.New("error not implemented")
	ErrFieldValidation          = errors.New("invalid field values")
	ErrRequestBodyParsingFailed = errors.New("request body parsing failed")
)

type FieldValidationError struct {
	ValidationErrors error
}

func (e *FieldValidationError) Error() string {
	if e.ValidationErrors == nil {
		return ErrFieldValidation.Error()
	}

	return e.ValidationErrors.Error()
}

func (e *FieldValidationError) Is(target error) bool {
	return target == ErrFieldValidation
}

func (e *FieldValidationError) HTTPResponse() apires.ErrorDetails {
	if e.ValidationErrors == nil {
		return apires.ErrorDetails{
			HTTPStatusCode: http.StatusBadRequest,
			Message:        pointer.New(e.Error()),
		}
	}

	return apires.ErrorDetails{
		HTTPStatusCode: http.StatusUnprocessableEntity,
		FieldErrors:    e.ValidationErrors,
	}
}
