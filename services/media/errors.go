package media

import "errors"

var (
	ErrInvalidMediaTypeValue       = errors.New("invalid media type value")
	ErrCreatingNewMediaItem        = errors.New("error creating new media item")
	ErrMediaIDGeneration           = errors.New("error generating media id")
	ErrInvalidNumberOfNewMediaItem = errors.New("invalid number of new media item")
	ErrInvalidMediaSize            = errors.New("invalid media size")
)
