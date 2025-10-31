package usermanagement

import (
	"errors"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/apires"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

var (
	ErrUserNotExistsInManagement   = errors.New("user does not exist in management")
	ErrUserGroupWithSameNameExists = errors.New("user group with same name exists")
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

type UserGroupWithSameNameExistsError struct{}

func (e *UserGroupWithSameNameExistsError) Error() string {
	return ErrUserGroupWithSameNameExists.Error()
}

func (e *UserGroupWithSameNameExistsError) Is(target error) bool {
	return target == ErrUserGroupWithSameNameExists
}

func (e *UserGroupWithSameNameExistsError) HTTPResponse() apires.ErrorDetails {
	return apires.ErrorDetails{
		HTTPStatusCode: http.StatusConflict,
		Message:        pointer.New(e.Error()),
	}
}
