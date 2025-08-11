package core

type ICache[T any] interface {
	Get(string) (T, bool)
	Set(string, T)
	Delete(string)
}

type ICacheClient interface {
	Get(string) any
	Set(string, any)
	Delete(string)
}
