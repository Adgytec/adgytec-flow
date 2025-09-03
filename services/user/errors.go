package user

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

var (
	ErrInvalidUserId = errors.New("invalid user id")
	ErrUserNotFound  = errors.New("user not found")
)

type InvalidUserIDError struct {
	InvalidUserID string
}

func (e *InvalidUserIDError) Error() string {
	return fmt.Sprintf("User ID: '%s', is not a valid user id.", e.InvalidUserID)
}

func (e *InvalidUserIDError) Is(target error) bool {
	return target == ErrInvalidUserId
}

func (e *InvalidUserIDError) HTTPResponse() core.ResponseHTTPError {
	return core.ResponseHTTPError{
		HTTPStatusCode: http.StatusBadRequest,
		Message:        pointer.New(e.Error()),
	}
}

type UserNotFoundError struct{}

func (e *UserNotFoundError) Error() string {
	return "User not found."
}

func (e *UserNotFoundError) Is(target error) bool {
	return target == ErrUserNotFound
}

func (e *UserNotFoundError) HTTPResponse() core.ResponseHTTPError {
	return core.ResponseHTTPError{
		HTTPStatusCode: http.StatusNotFound,
		Message:        pointer.New(e.Error()),
	}
}
