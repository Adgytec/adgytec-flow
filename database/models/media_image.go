package models

type ImageQueryType struct {
	OriginalImage string  `json:"originalImage"`
	Size          int     `json:"size"`
	Status        string  `json:"status"`
	Thumbnail     *string `json:"thumbnail,omitempty"`
	Small         *string `json:"small,omitempty"`
	Medium        *string `json:"medium,omitempty"`
	Large         *string `json:"large,omitempty"`
	ExtraLarge    *string `json:"extraLarge,omitempty"`
}

type VideoQueryType struct {
	OriginalVideo    string  `json:"originalVideo"`
	Size             int     `json:"size"`
	Status           string  `json:"status"`
	Thumbnail        *string `json:"thumbnail,omitempty"`
	AdaptiveManifest *string `json:"adaptiveManifest,omitempty"`
	Preview          *string `json:"preview,omitempty"`
}
