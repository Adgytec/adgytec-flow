package storage

import (
	"errors"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/apires"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

var (
	ErrInvalidS3ConfigValue = errors.New("invalid s3 config value")
	ErrInvalidPartNumbers   = errors.New("invalid part number")
)

type InvalidPartNumbersError struct{}

func (e *InvalidPartNumbersError) Error() string {
	return ErrInvalidPartNumbers.Error()
}

func (e *InvalidPartNumbersError) Is(target error) bool {
	return target == ErrInvalidPartNumbers
}

func (e *InvalidPartNumbersError) HTTPResponse() apires.ErrorDetails {
	return apires.ErrorDetails{
		HTTPStatusCode: http.StatusBadRequest,
		Message:        pointer.New(e.Error()),
	}
}
