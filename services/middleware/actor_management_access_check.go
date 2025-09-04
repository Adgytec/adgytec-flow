package app_middleware

import (
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/actor"
)

func (pc *appMiddlewarePC) ActorManagementAccessCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func (s *appMiddleware) actorManagementAccessCheck(currentActor actor.ActorDetails) error {
	return nil
}
