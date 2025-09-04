package app_errors

import (
	"errors"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/apires"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

var (
	ErrRequestDecode     = errors.New("request decoding failed")
	ErrRequestValidation = errors.New("request validation failed")
)

type RequestDecodeError struct {
	Status  int
	Message string
}

func (e *RequestDecodeError) Error() string {
	return e.Message
}

func (e *RequestDecodeError) Is(target error) bool {
	return target == ErrRequestDecode
}

func (e *RequestDecodeError) HTTPResponse() apires.ErrorDetails {
	return apires.ErrorDetails{
		HTTPStatusCode: e.Status,
		Message:        pointer.New(e.Error()),
	}
}

type RequestValidationError struct {
	FieldErrors map[string]string
}

func (e *RequestValidationError) Error() string {
	return "invalid request body"
}

func (e *RequestValidationError) Is(target error) bool {
	return target == ErrRequestValidation
}

func (e *RequestValidationError) HTTPResponse() apires.ErrorDetails {
	return apires.ErrorDetails{
		HTTPStatusCode: http.StatusUnprocessableEntity,
		FieldErrors:    &e.FieldErrors,
	}
}
