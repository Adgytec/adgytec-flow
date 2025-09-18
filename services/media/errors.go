package media

import "errors"

var (
	ErrInvalidMediaTypeValue = errors.New("invalid media type value")
	ErrCreatingNewMediaItem  = errors.New("error creating new media item")
)
