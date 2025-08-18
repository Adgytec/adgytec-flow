package app_errors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
)

var (
	ErrPermissionDenied      = errors.New("permission denied")
	ErrPermissionCheckFailed = errors.New("permission check failed")
)

type PermissionDeniedError struct {
	Action string
}

func (e *PermissionDeniedError) Error() string {
	return fmt.Sprintf("Permission denied for action: '%s'.", e.Action)
}

func (e *PermissionDeniedError) Is(target error) bool {
	return target == ErrPermissionDenied
}

func (e *PermissionDeniedError) HTTPResponse() core.ResponseHTTPError {
	return core.ResponseHTTPError{
		HTTPStatusCode: http.StatusForbidden,
		Message:        helpers.ValuePtr(e.Error()),
	}
}

type PermissionCheckFailedError struct {
	cause error
}

func (e *PermissionCheckFailedError) Error() string {
	return "Permission check failed."
}

func (e *PermissionCheckFailedError) Is(target error) bool {
	return target == ErrPermissionCheckFailed
}

func (e *PermissionCheckFailedError) Unwrap() error {
	return e.cause
}

func (e *PermissionCheckFailedError) HTTPResponse() core.ResponseHTTPError {
	// TODO: handle status code based on e.cause
	return core.ResponseHTTPError{
		HTTPStatusCode: http.StatusInternalServerError,
	}
}
