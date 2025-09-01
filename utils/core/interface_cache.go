package core

type ICache[T any] interface {
	Get(string, func() (T, error)) (T, error)
	Delete(string)
}

type ICacheClient interface {
	Get(string) (any, bool)
	Set(string, any)
	Delete(string)
}
