package serializer

import (
	"encoding/json"
)

type jsonSerializer[T any] struct{}

func (j *jsonSerializer[T]) Encode(data T) ([]byte, error) {
	return json.Marshal(data)
}

// value should be pointer
func (j *jsonSerializer[T]) Decode(data []byte) (T, error) {
	var value T
	jsonSerializerErr := json.Unmarshal(data, &value)
	return value, jsonSerializerErr
}

func NewJSONSerializer[T any]() core.Serializer[T] {
	return &jsonSerializer[T]{	}
}
