package user

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func (s *userService) removeUserSocialLink(ctx context.Context, userID, resourceID uuid.UUID) error {
	return nil
}

func (m *userServiceMux) removeUserSelfSocialLink(w http.ResponseWriter, r *http.Request) {}

func (m *userServiceMux) removeUserSocialLink(w http.ResponseWriter, r *http.Request) {}
