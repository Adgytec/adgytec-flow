package apikey

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/apires"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

var (
	ErrInvalidApiKey = errors.New("invalid api key")
)

type InvalidApiKeyError struct {
	apiKey string
}

func (e *InvalidApiKeyError) Error() string {
	return fmt.Sprintf("got invalid api key value: %s", e.apiKey)
}

func (e *InvalidApiKeyError) Is(target error) bool {
	return target == ErrInvalidApiKey
}

func (e *InvalidApiKeyError) HTTPResponse() apires.ErrorDetails {
	return apires.ErrorDetails{
		HTTPStatusCode: http.StatusBadRequest,
		Message:        pointer.New(ErrInvalidApiKey.Error()),
	}
}
