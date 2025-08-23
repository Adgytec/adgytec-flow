package app_middleware

import (
	"net/http"

	"github.com/google/uuid"
)

func (pc *appMiddlewarePC) ActorOrganizationManagementAccessCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// ensure adding actor is user middleware
func (s *appMiddleware) actorOrganizationManagementAccessCheck(actorID uuid.UUID, orgID uuid.UUID) error {
	return nil
}
