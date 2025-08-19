package app_errors

func valuePtr[T any](s T) *T {
	return &s
}
