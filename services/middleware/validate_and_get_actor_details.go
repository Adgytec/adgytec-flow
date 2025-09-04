package app_middleware

import (
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/actor"
)

func (pc *appMiddlewarePC) ValidateAndGetActorDetailsFromHttpRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func (s *appMiddleware) validateAndGetActorDetails(r *http.Request) (actor.ActorDetails, error) {
	return actor.ActorDetails{}, nil
}
