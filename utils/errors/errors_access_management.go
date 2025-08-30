package app_errors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

var (
	ErrPermissionDenied           = errors.New("permission denied")
	ErrPermissionResolutionFailed = errors.New("permission resolution failed")
)

type PermissionDeniedError struct {
	MissingPermission string
}

func (e *PermissionDeniedError) Error() string {
	return fmt.Sprintf("Permission denied. Missing required permissison: '%s'.", e.MissingPermission)
}

func (e *PermissionDeniedError) Is(target error) bool {
	return target == ErrPermissionDenied
}

func (e *PermissionDeniedError) HTTPResponse() core.ResponseHTTPError {
	return core.ResponseHTTPError{
		HTTPStatusCode: http.StatusForbidden,
		Message:        valuePtr(e.Error()),
	}
}

type PermissionResolutionFailedError struct {
	cause error
}

func (e *PermissionResolutionFailedError) Error() string {
	return "Permission resolution failed."
}

func (e *PermissionResolutionFailedError) Is(target error) bool {
	return target == ErrPermissionResolutionFailed
}

func (e *PermissionResolutionFailedError) Unwrap() error {
	return e.cause
}

func (e *PermissionResolutionFailedError) HTTPResponse() core.ResponseHTTPError {
	// TODO: handle status code based on e.cause
	return core.ResponseHTTPError{
		HTTPStatusCode: http.StatusInternalServerError,
	}
}
