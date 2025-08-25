package app_middleware

import (
	"net/http"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
)

func (pc *appMiddlewarePC) EnsureActorTypeUserOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func (s *appMiddleware) ensureActorTypeUserOnly(actorType db_actions.GlobalActorType) error {
	return nil
}
