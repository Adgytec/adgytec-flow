package models

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type GlobalUser struct {
	ID             uuid.UUID          `json:"id"`
	Email          string             `json:"email"`
	Name           string             `json:"name"`
	About          *string            `json:"about,omitempty"`
	DateOfBirth    pgtype.Date        `json:"dateOfBirth"`
	CreatedAt      pgtype.Timestamptz `json:"createdAt"`
	LastAccessed   pgtype.Timestamptz `json:"lastAccessed"`
	ProfilePicture *ImageQueryType    `json:"profilePicture,omitempty"`
}
