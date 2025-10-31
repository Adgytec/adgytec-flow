package models

import (
	"time"

	"github.com/google/uuid"
)

type UserGroup struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UserCount   int64     `json:"userCount"`
}

func (u UserGroup) GetCreatedAt() time.Time {
	return u.CreatedAt
}
