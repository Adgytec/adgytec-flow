package app_middleware

import (
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/db"
)

func (pc *appMiddlewarePC) EnsureActorTypeUserOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func (s *appMiddleware) ensureActorTypeUserOnly(actorType db.GlobalActorType) error {
	return nil
}
