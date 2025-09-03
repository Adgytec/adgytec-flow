package core

type Storage interface {
	GetPresignUploadUrl(string) (string, error)
}
