package user

import (
	"context"
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/google/uuid"
)

func (s *userService) updateUserSocialLink(ctx context.Context, userID, resourceID uuid.UUID) (*db.GlobalUserSocialLinks, error) {
	return nil, nil
}

func (m *userServiceMux) updateUserSelfSocialLink(w http.ResponseWriter, r *http.Request) {}
