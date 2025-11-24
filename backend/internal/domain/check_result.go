package domain

import (
	"time"

	"github.com/google/uuid"
)

type CheckResult struct {
	ID, MonitorID uuid.UUID
	StatusCode    int
	ResponseTime  int // ms
	Success       bool
	ErrorMessage  string
	CreatedAt     time.Time
}
