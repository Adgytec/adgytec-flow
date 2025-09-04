package storage

type Storage interface {
	GetPresignUploadUrl(string) (string, error)
}
