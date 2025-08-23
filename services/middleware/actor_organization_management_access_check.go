package app_middleware

import (
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/google/uuid"
)

func (pc *appMiddlewarePC) ActorOrganizationManagementAccessCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func (s *appMiddleware) actorOrganizationManagementAccessCheck(actor core.ActorDetails, orgID uuid.UUID) error {
	return nil
}
