package serializer

import (
	"encoding/json"
)

// used for primitive types
type jsonSerializer[T any] struct{}

func (j *jsonSerializer[T]) Encode(data T) ([]byte, error) {
	return json.Marshal(data)
}

func (j *jsonSerializer[T]) Decode(data []byte) (T, error) {
	var value T
	decodingErr := json.Unmarshal(data, &value)
	return value, decodingErr
}

func NewJSONSerializer[T any]() Serializer[T] {
	return &jsonSerializer[T]{}
}
