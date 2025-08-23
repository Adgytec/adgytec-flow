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

// before this middleware ensure adding middleware to check actor type is user
func (s *appMiddleware) actorOrganizationManagementAccessCheck(actorID uuid.UUID, orgID uuid.UUID) error {
	return nil
}
