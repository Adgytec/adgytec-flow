package serializer

import (
	"bytes"
	"encoding/gob"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

// used for custom struct and large payloads of data
type gobSerializer[T any] struct{}

func (g *gobSerializer[T]) Encode(data T) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	encodingErr := encoder.Encode(data)
	if encodingErr != nil {
		return nil, encodingErr
	}
	return buffer.Bytes(), nil
}

func (g *gobSerializer[T]) Decode(data []byte) (T, error) {
	var value T

	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)

	decodingErr := decoder.Decode(&value)
	return value, decodingErr
}

// NewGobSerializer should only be used for custom struct type
// it will throw runtime error for primitive errors during registration
func NewGobSerializer[T any]() core.Serializer[T] {
	var zero T
	gob.Register(&zero)

	return &gobSerializer[T]{}
}
