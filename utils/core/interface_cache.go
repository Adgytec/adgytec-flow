package core

type ICache[T any] interface {
	Get(string, func() (T, error)) (T, error)
	Delete(string)
}

type ICacheClient interface {
	Get(string) ([]byte, bool)
	Set(string, []byte)
	Delete(string)
}
