package reqparams

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/apires"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

var (
	ErrInvalidID          = errors.New("invalid id")
	ErrInvalidUserID      = errors.New("invalid user id")
	ErrInvalidUserGroupID = errors.New("invalid user group id")
)

type InvalidIDError struct {
	IDType    idType
	InvalidID string
}

func (e *InvalidIDError) Error() string {
	return fmt.Sprintf("%s: '%s', is not a valid id.", e.IDType, e.InvalidID)
}

func (e *InvalidIDError) Is(target error) bool {
	return target == ErrInvalidID
}

func (e *InvalidIDError) HTTPResponse() apires.ErrorDetails {
	return apires.ErrorDetails{
		HTTPStatusCode: http.StatusBadRequest,
		Message:        pointer.New(e.Error()),
	}
}
