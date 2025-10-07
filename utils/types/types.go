package types

import (
	"bytes"
	"encoding/json"
)

type Nullable[T any] struct {
	Value T
	Set   bool // if set is false means missing this field in req body
	Valid bool // if valid is false means value is considered nil
}

func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	n.Set = true

	// check for explicit null
	if bytes.Equal(data, []byte("null")) {
		n.Valid = false
		return nil
	}

	n.Valid = true
	return json.Unmarshal(data, &n.Value)
}

func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if !n.Set || !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Value)
}
