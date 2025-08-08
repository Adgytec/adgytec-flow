package core

type ICache[T any] interface {
	Get(string) T
	Set(string, T) error
	Delete(string)
}
