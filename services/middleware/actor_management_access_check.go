package app_middleware

import (
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

func (pc *appMiddlewarePC) ActorManagementAccessCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func (s *appMiddleware) actorManagementAccessCheck(actor core.ActorDetails) error {
	return nil
}
