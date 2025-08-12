package models

type MediaImageVariants struct {
	Thumbnail  string `json:"thumbnail"`
	Small      string `json:"small"`
	Medium     string `json:"medium"`
	Large      string `json:"large"`
	ExtraLarge string `json:"extraLarge"`
}

type ImageQueryType struct {
	OriginalImage string             `json:"originalImage"`
	Size          int                `json:"size"`
	Status        string             `json:"status"`
	Variants      MediaImageVariants `json:"variants"`
}

type VideoQueryType struct {
	OriginalVideo    string `json:"originalVideo"`
	Size             int    `json:"size"`
	Status           string `json:"status"`
	Thumbnail        string `json:"thumbnail"`
	AdaptiveManifest string `json:"adaptiveManifest"`
	Preview          string `json:"preview"`
}
