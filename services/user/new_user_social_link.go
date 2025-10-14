package user

import (
	"context"
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/google/uuid"
)

func (s *userService) newUserSocialLink(ctx context.Context, userID uuid.UUID) (*db.GlobalUserSocialLinks, error) {
	return nil, nil
}

func (m *userServiceMux) newUserSelfSocialLink(w http.ResponseWriter, r *http.Request) {}
