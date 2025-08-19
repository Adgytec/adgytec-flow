package app_errors

import "errors"

var (
	ErrInvalidActorDetails = errors.New("invalid actor details in context value")
)
