package appmiddleware

import (
	"net/http"
	"time"

	"github.com/Adgytec/adgytec-flow/utils/logger"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (pc *appMiddlewarePC) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// create or get request id
		reqID := r.Header.Get("X-Request-ID")
		if reqID == "" {
			reqID = uuid.New().String()
		}

		// wrap response writer to capture status code
		rw := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		// create logger with context
		reqLogger := log.With().Str("request_id", reqID).Logger()
		reqCtxWithLogger := logger.AddLoggerToContext(&reqLogger, r.Context())
		r = r.WithContext(reqCtxWithLogger)

		// log request details
		defer func() {
			reqLogger.Info().
				Str("method", r.Method).
				Str("url", r.URL.RequestURI()).
				Str("path", r.URL.Path).
				Int("status", rw.Status()).
				Int("bytes_written", rw.BytesWritten()).
				Str("user_agent", r.UserAgent()).
				Str("remote_ip", r.RemoteAddr).
				Dur("elapsed_ms", time.Since(start)).
				Msg("incoming request")
		}()

		next.ServeHTTP(rw, r)
	})
}
