package serializer

type gobSerializer struct{}

func (gob *gobSerializer) encode(data any) ([]byte, error) {
	return nil, nil
}

func (gob *gobSerializer) decode(data []byte) (any, error) {
	return nil, nil
}

func createGobSerializer() iSerializer {
	return &gobSerializer{}
}
