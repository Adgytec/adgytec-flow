package core

type ICDN interface {
	GetSignedUrl(*string) *string
}
