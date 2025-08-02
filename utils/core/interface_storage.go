package core

type IStorage interface {
	GetPresignUploadUrl(string) (string, error)
}
