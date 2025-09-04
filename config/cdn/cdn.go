package cdn

type CDN interface {
	GetSignedUrl(*string) *string
}
