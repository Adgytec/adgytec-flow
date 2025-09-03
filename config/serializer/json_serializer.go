package serializer

import (
	"encoding/json"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

type jsonSerializer[T any] struct{}

func (j *jsonSerializer[T]) Encode(data T) ([]byte, error) {
	return json.Marshal(data)
}

// value should be pointer
func (j *jsonSerializer[T]) Decode(data []byte) (T, error) {
	var value T
	decodingErr := json.Unmarshal(data, &value)
	return value, decodingErr
}

func NewJSONSerializer[T any]() core.Serializer[T] {
	return &jsonSerializer[T]{}
}
