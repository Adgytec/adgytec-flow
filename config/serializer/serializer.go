package serializer

import (
	"github.com/Adgytec/adgytec-flow/utils/core"
)

// iSerializer defines the internal serializers like json-serializer, gob-serializer methods
type iSerializer interface {
	encode(any) ([]byte, error)
	decode([]byte, any) error
}

type serializer[T any] struct {
	serializer iSerializer
}

func (s *serializer[T]) Encode(data T) ([]byte, error) {
	return s.serializer.encode(data)
}

func (s *serializer[T]) Decode(data []byte) (T, error) {
	var value T

	serializeErr := s.serializer.decode(data, &value)
	return value, serializeErr
}

func CreateSerializer[T any]() core.ISerializer[T] {
	return &serializer[T]{
		serializer: createJsonSerializer(),
	}
}
