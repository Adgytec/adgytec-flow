package usermanagement

import (
	"errors"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/apires"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

var (
	ErrUserNotExistsInManagement = errors.New("user not exists in management")
)

type UserNotExistsInManagementError struct{}

func (e *UserNotExistsInManagementError) Error() string {
	return ErrUserNotExistsInManagement.Error()
}

func (e *UserNotExistsInManagementError) Is(target error) bool {
	return target == ErrUserNotExistsInManagement
}

func (e *UserNotExistsInManagementError) HTTPResponse() apires.ErrorDetails {
	return apires.ErrorDetails{
		HTTPStatusCode: http.StatusNotFound,
		Message:        pointer.New(e.Error()),
	}
}
