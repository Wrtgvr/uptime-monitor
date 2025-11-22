package domain

import (
	"time"

	"github.com/google/uuid"
)

type Monitor struct {
	ID, UserID             uuid.UUID
	Name, URL, Method      string
	CheckInterval, Timeout int // sec
	ExpectedCodeRange      int // 100, 200 .. 500
	CreatedAt, NextCheckAt time.Time
}

type MonitorCreate struct {
	Name, URL, Method      string
	CheckInterval, Timeout int // sec
	ExpectedCodeRange      int // 100, 200 .. 500
}

type MonitorUpdate struct {
	ID                     uuid.UUID
	Name, Method           string
	CheckInterval, Timeout int // sec
	ExpectedCodeRange      int // 100, 200 .. 500
}
