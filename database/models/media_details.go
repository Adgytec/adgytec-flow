package models

import "github.com/google/uuid"

type ImageDetails struct {
	MediaID       uuid.UUID `json:"-"` // used for internal purposes
	OriginalImage *string   `json:"originalImage"`
	Size          *int64    `json:"size"`
	Status        *string   `json:"status"`
	Thumbnail     *string   `json:"thumbnail"`
	Small         *string   `json:"small"`
	Medium        *string   `json:"medium"`
	Large         *string   `json:"large"`
	ExtraLarge    *string   `json:"extraLarge"`
}

type VideoDetails struct {
	MediaID          uuid.UUID `json:"-"` // used for internal purposes
	OriginalVideo    *string   `json:"originalVideo"`
	Size             *int64    `json:"size"`
	Status           *string   `json:"status"`
	Thumbnail        *string   `json:"thumbnail"`
	AdaptiveManifest *string   `json:"adaptiveManifest"`
	Preview          *string   `json:"preview"`
}
