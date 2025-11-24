package domain

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	JTI, TokenHash                  string
	UserID                          uuid.UUID
	ExpiresAt, UpdatedAt, CreatedAt time.Time
	Revoked                         bool
}

type RefreshTokenCreate struct {
	Token     string
	UserID    uuid.UUID
	ExpiresAt time.Time
	Revoked   bool
}
