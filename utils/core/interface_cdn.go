package core

type ICDN interface {
	GetSignedUrl(string) (string, error)
	GetSignedUrls([]string) ([]string, error)
}
