package serializer

type Serializer[T any] interface {
	Encode(data T) ([]byte, error)
	Decode(data []byte) (T, error)
}
