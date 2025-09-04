package actor

import "errors"

var (
	ErrInvalidActorID   = errors.New("invalid actor id")
	ErrInvalidActorType = errors.New("invalid actor type")
)
