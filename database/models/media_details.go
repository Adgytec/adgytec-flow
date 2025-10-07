package models

type ImageDetails struct {
	OriginalImage *string `json:"originalImage"`
	Size          *uint64 `json:"size"`
	Status        *string `json:"status"`
	Thumbnail     *string `json:"thumbnail"`
	Small         *string `json:"small"`
	Medium        *string `json:"medium"`
	Large         *string `json:"large"`
	ExtraLarge    *string `json:"extraLarge"`
}

type VideoDetails struct {
	OriginalVideo    *string `json:"originalVideo"`
	Size             *uint64 `json:"size"`
	Status           *string `json:"status"`
	Thumbnail        *string `json:"thumbnail"`
	AdaptiveManifest *string `json:"adaptiveManifest"`
	Preview          *string `json:"preview"`
}
