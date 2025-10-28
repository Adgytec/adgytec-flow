package usermanagement

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func (s *userManagementService) removeUser(ctx context.Context, userID uuid.UUID) error {
	return nil
}

func (m *serviceMux) removeUser(w http.ResponseWriter, r *http.Request) {}
