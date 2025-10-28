package user

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/apires"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

var (
	ErrUserNotFound                           = errors.New("user not found")
	ErrNameLength                             = errors.New("User name must be between 3 and 100 characters long.")
	ErrAboutLength                            = errors.New("User about must be between 8 and 1024 characters long.")
	ErrInvalidDateOfBirth                     = errors.New("Invalid date of birth")
	ErrUserSocialPlatformDetailsAlreadyExists = errors.New("user social platform details already exists")
	ErrSocialLinkNotFound                     = errors.New("social link not found")
	ErrInvalidSocialLinkID                    = errors.New("invalid social link id")
)

type UserNotFoundError struct{}

func (e *UserNotFoundError) Error() string {
	return "User not found."
}

func (e *UserNotFoundError) Is(target error) bool {
	return target == ErrUserNotFound
}

func (e *UserNotFoundError) HTTPResponse() apires.ErrorDetails {
	return apires.ErrorDetails{
		HTTPStatusCode: http.StatusNotFound,
		Message:        pointer.New(e.Error()),
	}
}

type UserSocialPlatformDetailsAlreadyExistsError struct {
	PlatformName string
}

func (e *UserSocialPlatformDetailsAlreadyExistsError) Error() string {
	return fmt.Sprintf("'%s' details already exists", e.PlatformName)
}

func (e *UserSocialPlatformDetailsAlreadyExistsError) Is(target error) bool {
	return target == ErrUserSocialPlatformDetailsAlreadyExists
}

func (e *UserSocialPlatformDetailsAlreadyExistsError) HTTPResponse() apires.ErrorDetails {
	return apires.ErrorDetails{
		HTTPStatusCode: http.StatusConflict,
		Message:        pointer.New(e.Error()),
	}
}

type SocialLinkNotFoundError struct{}

func (e *SocialLinkNotFoundError) Error() string {
	return "Social link not found."
}

func (e *SocialLinkNotFoundError) Is(target error) bool {
	return target == ErrSocialLinkNotFound
}

func (e *SocialLinkNotFoundError) HTTPResponse() apires.ErrorDetails {
	return apires.ErrorDetails{
		HTTPStatusCode: http.StatusNotFound,
		Message:        pointer.New(e.Error()),
	}
}

type InvalidSocialLinkIDError struct {
	InvalidSocialLinkID string
}

func (e *InvalidSocialLinkIDError) Error() string {
	return fmt.Sprintf("ID: '%s', is not a valid social link id.", e.InvalidSocialLinkID)
}

func (e *InvalidSocialLinkIDError) Is(target error) bool {
	return target == ErrInvalidSocialLinkID
}

func (e *InvalidSocialLinkIDError) HTTPResponse() apires.ErrorDetails {
	return apires.ErrorDetails{
		HTTPStatusCode: http.StatusBadRequest,
		Message:        pointer.New(e.Error()),
	}
}
