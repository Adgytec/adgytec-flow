package core

type Cache[T any] interface {
	Get(string, func() (T, error)) (T, error)
	Delete(string)
}

type CacheClient interface {
	Get(string) ([]byte, bool)
	Set(string, []byte)
	Delete(string)
}
