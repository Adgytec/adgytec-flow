package appmiddleware

import "net/http"

// responseWriter captures statusCode for the response
type responseWriter struct {
	statusCode int
	http.ResponseWriter
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.statusCode = status
	rw.ResponseWriter.WriteHeader(status)
}
