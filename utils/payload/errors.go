package payload

import (
	"errors"

	"github.com/Adgytec/adgytec-flow/utils/apires"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

var (
	ErrRequestDecode = errors.New("request decoding failed")
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
