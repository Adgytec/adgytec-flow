package helpers

func ValuePtr[T any](s T) *T {
	return &s
}
