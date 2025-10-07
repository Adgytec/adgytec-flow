package models

import (
	"time"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type GlobalUser struct {
	ID             uuid.UUID           `json:"id"`
	Email          string              `json:"email"`
	Name           string              `json:"name"`
	About          *string             `json:"about"`
	DateOfBirth    pgtype.Date         `json:"dateOfBirth"`
	CreatedAt      time.Time           `json:"createdAt"`
	ProfilePicture *ImageDetails       `json:"profilePicture"`
	Status         db.GlobalUserStatus `json:"status"`
}

func (u GlobalUser) GetCreatedAt() time.Time {
	return u.CreatedAt
}
