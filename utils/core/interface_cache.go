package core

type ICache[T any] interface {
	Get(string) T
	Set(string, T) error
	Delete(string)
}

type ICacheClient interface {
	Get(string) any
	Set(string, any) error
	Delete(string)
}
