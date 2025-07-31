package payload

import "net/http"

func DecodeRequest[T any](w http.ResponseWriter, r *http.Request) error {
	// 8 kilo byte
	r.Body = http.MaxBytesReader(w, r.Body, 1<<13)
	return nil
}
