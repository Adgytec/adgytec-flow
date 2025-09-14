package models

type ImageDetails struct {
	OriginalImage *string `json:"originalImage"`
	Size          *int64  `json:"size"`
	Status        *string `json:"status,omitempty"`
	Thumbnail     *string `json:"thumbnail,omitempty"`
	Small         *string `json:"small,omitempty"`
	Medium        *string `json:"medium,omitempty"`
	Large         *string `json:"large,omitempty"`
	ExtraLarge    *string `json:"extraLarge,omitempty"`
}

type VideoDetails struct {
	OriginalVideo    *string `json:"originalVideo"`
	Size             *int64  `json:"size"`
	Status           *string `json:"status,omitempty"`
	Thumbnail        *string `json:"thumbnail,omitempty"`
	AdaptiveManifest *string `json:"adaptiveManifest,omitempty"`
	Preview          *string `json:"preview,omitempty"`
}
