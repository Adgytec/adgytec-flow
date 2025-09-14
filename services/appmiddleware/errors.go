package appmiddleware

import (
	"errors"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/apires"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
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

func (e *UserStatusBadError) HTTPResponse() apires.ErrorDetails {
	return apires.ErrorDetails{
		HTTPStatusCode: http.StatusUnauthorized,
		Message:        pointer.New(http.StatusText(http.StatusUnauthorized)),
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

func (e *OrganizationStatusBadError) HTTPResponse() apires.ErrorDetails {
	return apires.ErrorDetails{
		HTTPStatusCode: http.StatusForbidden,
		Message:        pointer.New("Action forbidden due to organization status."),
	}
}

type NoAccessError struct{}

func (e *NoAccessError) Error() string {
	return "no access"
}

func (e *NoAccessError) Is(target error) bool {
	return target == ErrNoAccess
}

func (e *NoAccessError) HTTPResponse() apires.ErrorDetails {
	return apires.ErrorDetails{
		HTTPStatusCode: http.StatusForbidden,
		Message:        pointer.New(http.StatusText(http.StatusForbidden)),
	}
}
