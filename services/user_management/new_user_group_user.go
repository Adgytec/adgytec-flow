package usermanagement

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func (s *userManagementService) newUserGroupUser(ctx context.Context, userData newUserData) (*uuid.UUID, error) {
	return nil, nil
}

func (m *serviceMux) newUserGroupUser(w http.ResponseWriter, r *http.Request) {}
