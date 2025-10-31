package usermanagement

import (
	"context"
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/models"
	"github.com/google/uuid"
)

func (s *userManagementService) getUserGroupByID(ctx context.Context, groupID uuid.UUID) (*models.UserGroup, error) {
	return nil, nil
}

func (m *serviceMux) getUserGroupByID(w http.ResponseWriter, r *http.Request) {}
