package app_errors

import (
	"fmt"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
)

var (
	ErrUserExists       = fmt.Errorf("user exists")
	ErrAuthActionFailed = fmt.Errorf("auth action failed")
)

type UserExistsError struct {
	username string
}

func (e *UserExistsError) Error() string {
	return fmt.Sprintf("User with username %s already exists.", e.username)
}

func (e *UserExistsError) Is(target error) bool {
	return target == ErrUserExists
}

func (e *UserExistsError) HTTPResponse() core.ResponseHTTPError {
	return core.ResponseHTTPError{
		HTTPStatusCode: http.StatusConflict,
		Message:        helpers.StringPtr(e.Error()),
	}
}

type AuthActionFailedError struct {
	username   string
	reason     string
	actionType core.AuthActionType
	cause      error
}

func (e *AuthActionFailedError) Error() string {
	return fmt.Sprintf("Auth action failed.\nAction Type: '%s' \nUsername: %s\nReason: %s", e.actionType, e.username, e.reason)
}

func (e *AuthActionFailedError) Is(target error) bool {
	return target == ErrAuthActionFailed
}

func (e *AuthActionFailedError) Unwrap() error {
	return e.cause
}

func (e *AuthActionFailedError) HTTPResponse() core.ResponseHTTPError {
	// TODO: handle status based on e.cause
	return core.ResponseHTTPError{
		HTTPStatusCode: http.StatusInternalServerError,
	}
}
