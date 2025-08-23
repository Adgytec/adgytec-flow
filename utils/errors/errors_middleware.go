package app_errors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

var (
	ErrUserStatusBad         = errors.New("bad user status")
	ErrOrganizationStatusBad = errors.New("bad organization status")
	ErrNoAccess              = errors.New("no access")
)

type UserStatusBadError struct{}

func (e *UserStatusBadError) Error() string {
	return "bad user status"
}

func (e *UserStatusBadError) Is(target error) bool {
	return target == ErrUserStatusBad
}

func (e *UserStatusBadError) HTTPResponse() core.ResponseHTTPError {
	return core.ResponseHTTPError{
		HTTPStatusCode: http.StatusUnauthorized,
		Message:        valuePtr(http.StatusText(http.StatusUnauthorized)),
	}
}

// TODO: when organization service is created add organization status to error
type OrganizationStatusBadError struct{}

func (e *OrganizationStatusBadError) Error() string {
	return "bad organization status"
}

func (e *OrganizationStatusBadError) Is(target error) bool {
	return target == ErrOrganizationStatusBad
}

func (e *OrganizationStatusBadError) HTTPResponse() core.ResponseHTTPError {
	// TODO: message will be based on organization status
	return core.ResponseHTTPError{
		HTTPStatusCode: http.StatusForbidden,
		Message:        valuePtr(fmt.Sprintf("Organization is currently: %s", "status")),
	}
}

type NoAccessError struct{}

func (e *NoAccessError) Error() string {
	return "bad organization status"
}

func (e *NoAccessError) Is(target error) bool {
	return target == ErrNoAccess
}

func (e *NoAccessError) HTTPResponse() core.ResponseHTTPError {
	// TODO: message will be based on organization status
	return core.ResponseHTTPError{
		HTTPStatusCode: http.StatusForbidden,
		Message:        valuePtr(http.StatusText(http.StatusForbidden)),
	}
}
