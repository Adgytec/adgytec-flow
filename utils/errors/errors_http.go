package app_errors

import (
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
)

type RequestDecodeError struct {
	Status  int
	Message string
}

func (e *RequestDecodeError) Error() string {
	return e.Message
}

func (e *RequestDecodeError) HTTPResponse() core.ResponseHTTPError {
	return core.ResponseHTTPError{
		HTTPStatusCode: e.Status,
		Message:        helpers.ValuePtr(e.Error()),
	}
}

type RequestValidationError struct {
	FieldErrors map[string]string
}

func (e *RequestValidationError) Error() string {
	return "invalid request body"
}

func (e *RequestValidationError) HTTPResponse() core.ResponseHTTPError {
	return core.ResponseHTTPError{
		HTTPStatusCode: http.StatusUnprocessableEntity,
		FieldErrors:    &e.FieldErrors,
	}
}
