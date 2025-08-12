package app_errors

import (
	"fmt"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
)

var (
	ErrInvalidUserId = fmt.Errorf("invalid-user-id")
	ErrUserNotFound  = fmt.Errorf("user-not-found")
)

type InvalidUserIdError struct {
	InvalidUserId string
}

func (e *InvalidUserIdError) Error() string {
	return fmt.Sprintf("User ID: '%s', is not a valid user id.", e.InvalidUserId)
}

func (e *InvalidUserIdError) Is(target error) bool {
	return target == ErrInvalidUserId
}

func (e *InvalidUserIdError) HTTPResponse() core.ResponseHTTPError {
	return core.ResponseHTTPError{
		HTTPStatusCode: http.StatusBadRequest,
		Message:        helpers.StringPtr(e.Error()),
	}
}

type UserNotFoundError struct {
}

func (e *UserNotFoundError) Error() string {
	return "User not found."
}

func (e *UserNotFoundError) Is(target error) bool {
	return target == ErrUserNotFound
}

func (e *UserNotFoundError) HTTPResponse() core.ResponseHTTPError {
	return core.ResponseHTTPError{
		HTTPStatusCode: http.StatusNotFound,
		Message:        helpers.StringPtr(e.Error()),
	}
}
