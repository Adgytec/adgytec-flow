package serializer

import (
	"bytes"
	"encoding/gob"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

type gobSerializer[T any] struct{}

func (j *gobSerializer[T]) Encode(data T) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	encodingErr := encoder.Encode(data)
	if encodingErr != nil {
		return nil, encodingErr
	}
	return buffer.Bytes(), nil
}

// value should be pointer
func (j *gobSerializer[T]) Decode(data []byte) (T, error) {
	var value T

	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)

	decodingErr := decoder.Decode(&value)
	return value, decodingErr
}

func NewGobSerializer[T any]() core.Serializer[T] {
	return &gobSerializer[T]{}
}
