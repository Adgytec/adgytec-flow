package serializer

import (
	"github.com/Adgytec/adgytec-flow/utils/core"
	app_errors "github.com/Adgytec/adgytec-flow/utils/errors"
)

// iSerializer defines the internal serializers like json-serializer, gob-serializer methods
type iSerializer interface {
	encode(any) ([]byte, error)
	decode([]byte) (any, error)
}

type serializer[T any] struct {
	serializer iSerializer
}

func (s *serializer[T]) Encode(data T) ([]byte, error) {
	return s.serializer.encode(data)
}

func (s *serializer[T]) Decode(data []byte) (T, error) {
	var zero T

	serilizedData, serializeErr := s.serializer.decode(data)
	if serializeErr != nil {
		return zero, serializeErr
	}

	val, typeOK := serilizedData.(T)
	if !typeOK {
		return zero, app_errors.ErrTypeCastingCacheValueFailed
	}

	return val, nil
}

func CreateSerializer[T any]() core.ISerializer[T] {
	return &serializer[T]{
		serializer: createGobSerializer(),
	}
}
