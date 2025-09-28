package payload

import (
	"errors"

	"github.com/Adgytec/adgytec-flow/utils/apires"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

var (
	ErrRequestBodyDecode = errors.New("request decoding failed")
)

type RequestBodyDecodeError struct {
	Status  int
	Message string
}

func (e *RequestBodyDecodeError) Error() string {
	return e.Message
}

func (e *RequestBodyDecodeError) Is(target error) bool {
	return target == ErrRequestBodyDecode
}

func (e *RequestBodyDecodeError) HTTPResponse() apires.ErrorDetails {
	return apires.ErrorDetails{
		HTTPStatusCode: e.Status,
		Message:        pointer.New(e.Error()),
	}
}
