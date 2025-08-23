package app_middleware

import (
	"net/http"

	"github.com/google/uuid"
)

func (pc *appMiddlewarePC) CheckOrganizationStatusAndRequestMethod(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func (s *appMiddleware) checkOrganizationStatusAndRequestMethod(orgID uuid.UUID, method string) error {
	return nil
}
