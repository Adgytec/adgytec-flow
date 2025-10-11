package appmiddleware

import (
	"net/http"
	"runtime/debug"

	"github.com/Adgytec/adgytec-flow/utils/logger"
)

func (pc *appMiddlewarePC) Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				// Get logger from context
				log := logger.GetLoggerFromContext(r.Context())
				log.Error().
					Interface("panic", rvr).
					Bytes("stack", debug.Stack()).
					Msg("request panicked")

				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
