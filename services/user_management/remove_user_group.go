package usermanagement

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func (s *userManagementService) removeUserGroup(ctx context.Context, groupID uuid.UUID) error {
	return nil
}

func (m *serviceMux) removeUserGroup(w http.ResponseWriter, r *http.Request) {}
