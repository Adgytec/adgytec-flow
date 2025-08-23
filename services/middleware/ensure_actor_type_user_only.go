package app_middleware

import (
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

func (pc *appMiddlewarePC) EnsureActorTypeUserOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func (s *appMiddleware) ensureActorTypeUserOnly(actorType core.ActorType) error {
	return nil
}
