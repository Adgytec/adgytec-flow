package user

import (
	"strings"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/google/uuid"
)

func (s *userService) getUserIDFromEmail(email string) uuid.UUID {
	return core.GetIDFromPayload([]byte(strings.TrimSpace(email)))
}

func (pc *userServicePC) GetUserIDFromEmail(email string) uuid.UUID {
	return pc.service.getUserIDFromEmail(email)
}
