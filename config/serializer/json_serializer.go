package serializer

import (
	"encoding/json"
)

type jsonSerializer struct{}

func (j *jsonSerializer) encode(data any) ([]byte, error) {
	return json.Marshal(data)
}

// value should be pointer
func (j *jsonSerializer) decode(data []byte, value any) error {
	return json.Unmarshal(data, value)
}

func createJsonSerializer() iSerializer {
	return &jsonSerializer{}
}
