package pointer

func New[T any](data T) *T {
	return &data
}
