package iam

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

var (
	ErrPermissionDenied              = errors.New("permission denied")
	ErrPermissionResolutionFailed    = errors.New("permission resolution failed")
	ErrUnknownPermissionType         = errors.New("unknown permission type")
	ErrMissingRequiredResourcesValue = errors.New("missing required resources value")
	ErrMissingPermissionsToCheck     = errors.New("missing permissions to check")
)

// PermissionDeniedError defines error used when permission is denied for reasons that doesn't involve external errors
// MissingPermission tells which permission is missing
// Reason gives more details about why permission is denied, if permission resolution failed before actually checking if permission is present
// like permission actor type and current actor type doesn't match
// Only one of the MissingPermission or Reason is used for final Error() message and Reason is given more priority
type PermissionDeniedError struct {
	MissingPermission string
	Reason            string
}

func (e *PermissionDeniedError) Error() string {
	if e.Reason != "" {
		return fmt.Sprintf("Permission denied: %s", e.Reason)
	}

	if e.MissingPermission != "" {

		return fmt.Sprintf("Permission denied: missing required permission '%s'", e.MissingPermission)
	}

	return ErrPermissionDenied.Error()
}

func (e *PermissionDeniedError) Is(target error) bool {
	return target == ErrPermissionDenied
}

func (e *PermissionDeniedError) HTTPResponse() core.ResponseHTTPError {
	return core.ResponseHTTPError{
		HTTPStatusCode: http.StatusForbidden,
		Message:        pointer.New(e.Error()),
	}
}

type PermissionResolutionFailedError struct {
	Cause error
}

func (e *PermissionResolutionFailedError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("Permission resolution failed: %v", e.Cause)
	}
	return ErrPermissionResolutionFailed.Error()
}

func (e *PermissionResolutionFailedError) Is(target error) bool {
	return target == ErrPermissionResolutionFailed
}

func (e *PermissionResolutionFailedError) Unwrap() error {
	return e.Cause
}

func (e *PermissionResolutionFailedError) HTTPResponse() core.ResponseHTTPError {
	// TODO: handle status code based on e.cause
	return core.ResponseHTTPError{
		HTTPStatusCode: http.StatusInternalServerError,
	}
}
