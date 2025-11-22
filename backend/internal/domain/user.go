package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID                        uuid.UUID
	Username, Email, PasswordHash string
	Plan                          string // Im not sure about implementing premium plan but field just be here
	CreatedAt, UpdatedAt          time.Time
	GrandfatheredLimits           *GrandfatheredLimits
}

type UserCreate struct {
	Username, Email, Password string
}

type UserUpdate struct {
	UserID                    uuid.UUID
	Username, Email, Password string
}
