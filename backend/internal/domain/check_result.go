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

type CheckResultCreate struct {
	MonitorID                uuid.UUID
	StatusCode, ResponseTime int
	Success                  bool
	ErrorMessage             string
}
