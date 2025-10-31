package reqparams

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/apires"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

var (
	ErrInvalidUserID      = errors.New("invalid user id")
	ErrInvalidUserGroupID = errors.New("invalid user group id")
)

type InvalidUserIDError struct {
	InvalidUserID string
}

func (e *InvalidUserIDError) Error() string {
	return fmt.Sprintf("User ID: '%s', is not a valid user id.", e.InvalidUserID)
}

func (e *InvalidUserIDError) Is(target error) bool {
	return target == ErrInvalidUserID
}

func (e *InvalidUserIDError) HTTPResponse() apires.ErrorDetails {
	return apires.ErrorDetails{
		HTTPStatusCode: http.StatusBadRequest,
		Message:        pointer.New(e.Error()),
	}
}

type InvalidUserGroupIDError struct {
	InvalidUserGroupID string
}

func (e *InvalidUserGroupIDError) Error() string {
	return fmt.Sprintf("User Group ID: '%s', is not a valid group id.", e.InvalidUserGroupID)
}

func (e *InvalidUserGroupIDError) Is(target error) bool {
	return target == ErrInvalidUserGroupID
}

func (e *InvalidUserGroupIDError) HTTPResponse() apires.ErrorDetails {
	return apires.ErrorDetails{
		HTTPStatusCode: http.StatusBadRequest,
		Message:        pointer.New(e.Error()),
	}
}
