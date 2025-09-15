package media

import "github.com/Adgytec/adgytec-flow/utils/core"

type NewMediaItemInput struct{}

// Validate() validates the input values
// implementation details will be added later
func (mediaItemInput *NewMediaItemInput) Validate() error {
	return core.ErrNotImplemented
}

// EnsureMediaItemIsImage() ensures the item that will be uploaded is image
// implementation will be added later
func (mediaItemInput *NewMediaItemInput) EnsureMediaItemIsImage() error {
	validationErr := mediaItemInput.Validate()
	if validationErr != nil {
		return validationErr
	}

	return core.ErrNotImplemented
}

// EnsureMediaItemIsVideo() ensures the item that will be uploaded is video
// implementation will be added later
func (mediaItemInput *NewMediaItemInput) EnsureMediaItemIsVideo() error {
	validationErr := mediaItemInput.Validate()
	if validationErr != nil {
		return validationErr
	}

	return core.ErrNotImplemented
}

type NewMediaItemOutput struct{}
