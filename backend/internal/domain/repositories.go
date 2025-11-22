package domain

import (
	"context"

	"github.com/google/uuid"
	errs "github.com/wrtgvr/uptime-monitor/internal/errors"
)

type UsersRepo interface {
	GetUser(context.Context, uuid.UUID) (*User, *errs.AppError)
	CreateUser(context.Context, *UserCreate) *errs.AppError
	UpdateUser(context.Context, *UserUpdate) *errs.AppError
	DeleteUser(context.Context, uuid.UUID) *errs.AppError
}

type MonitorsRepo interface {
	GetMonitor(context.Context, uuid.UUID) *errs.AppError
	CreateMonitor(context.Context, *MonitorCreate) *errs.AppError
	UpdateMonitor(context.Context, *MonitorUpdate) *errs.AppError
	DeleteMonitor(context.Context, uuid.UUID) *errs.AppError
}

type CheckResultsRepo interface {
	GetCheckResult(context.Context, uuid.UUID) *errs.AppError
	CreateCheckResult(context.Context, *CheckResultCreate) *errs.AppError
	DeleteCheckResult(context.Context, uuid.UUID) *errs.AppError
}
