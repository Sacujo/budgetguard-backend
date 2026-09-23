package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Email        string
	PasswordHash string
	RevisionCtr  int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
