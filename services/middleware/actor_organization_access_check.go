package app_middleware

import (
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/actor"
	"github.com/google/uuid"
)

func (pc *appMiddlewarePC) ActorOrganizationAccessCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func (s *appMiddleware) actorOrganizationAccessCheck(currentActor actor.ActorDetails, orgID uuid.UUID) error {
	return nil
}
