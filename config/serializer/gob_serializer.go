package serializer

import (
	"bytes"
	"encoding/gob"
)

type gobSerializer struct{}

func (g *gobSerializer) encode(data any) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	encodingErr := encoder.Encode(data)
	if encodingErr != nil {
		return nil, encodingErr
	}

	return buffer.Bytes(), nil
}

// value should be pointer
func (g *gobSerializer) decode(data []byte, value any) error {
	buf := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buf)

	return decoder.Decode(value)
}

func createGobSerializer() iSerializer {
	return &gobSerializer{}
}
