package models

type MediaImageVariants struct {
	Thumbnail  string `json:"thumbnail"`
	Small      string `json:"small"`
	Medium     string `json:"medium"`
	Large      string `json:"large"`
	ExtraLarge string `json:"extraLarge"`
}
