package core

type CDN interface {
	GetSignedUrl(*string) *string
}
