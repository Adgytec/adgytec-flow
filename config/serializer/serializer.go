package serializer

// serializer defines the internal serializers like json-serializer, gob-serializer methods
type serializer interface {
	encode(any) ([]byte, error)
	decode([]byte, any) error
}

type defaultSerializer[T any] struct {
	serializer serializer
}

func (s *defaultSerializer[T]) Encode(data T) ([]byte, error) {
	return s.serializer.encode(data)
}

func (s *defaultSerializer[T]) Decode(data []byte) (T, error) {
	var value T

	serializeErr := s.serializer.decode(data, &value)
	return value, serializeErr
}
