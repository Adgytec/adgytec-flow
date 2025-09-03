package app_errors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

var (
	ErrUserExists         = errors.New("user exists")
	ErrAuthActionFailed   = errors.New("auth action failed")
	ErrInvalidAccessToken = errors.New("invalid access token")
	ErrInvalidAPIKey      = errors.New("invalid api key")
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
		Message:        pointer.New(e.Error()),
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

type InvalidAccessTokenError struct{}

func (e *InvalidAccessTokenError) Error() string {
	return "Invalid access token."
}

func (e *InvalidAccessTokenError) Is(target error) bool {
	return target == ErrInvalidAccessToken
}

func (e *InvalidAccessTokenError) HTTPResponse() core.ResponseHTTPError {
	return core.ResponseHTTPError{
		HTTPStatusCode: http.StatusBadRequest,
		Message:        pointer.New(e.Error()),
	}
}

type InvalidAPIKeyError struct{}

func (e *InvalidAPIKeyError) Error() string {
	return "Invalid API Key."
}

func (e *InvalidAPIKeyError) Is(target error) bool {
	return target == ErrInvalidAPIKey
}

func (e *InvalidAPIKeyError) HTTPResponse() core.ResponseHTTPError {
	return core.ResponseHTTPError{
		HTTPStatusCode: http.StatusBadRequest,
		Message:        pointer.New(e.Error()),
	}
}
