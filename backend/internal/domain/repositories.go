package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	errs "github.com/wrtgvr/uptime-monitor/internal/errors"
)

//! Repositorires CreateModel functions accept full Model instead of ModelCreate (e.g. User instead of UserCreate)
//! ModelCreate srtucts are used in service layer which create full Model structs to pass it to repository layer

type UserRepo interface {
	GetUserByID(context.Context, uuid.UUID) (*User, *errs.AppError)
	GetUserByEmail(context.Context, string) (*User, *errs.AppError)
	CreateUser(context.Context, *User) *errs.AppError
	UpdateUser(context.Context, *User) *errs.AppError
	DeleteUser(context.Context, uuid.UUID) *errs.AppError
}

type MonitorRepo interface {
	GetUserMonitors(context.Context, uuid.UUID) ([]*Monitor, *errs.AppError)
	GetMonitor(context.Context, uuid.UUID) (*Monitor, *errs.AppError)
	CreateMonitor(context.Context, *Monitor) *errs.AppError
	UpdateMonitor(context.Context, *Monitor) *errs.AppError
	DeleteMonitor(context.Context, uuid.UUID) *errs.AppError
}

type CheckResultRepo interface {
	GetLastMonitorResults(ctx context.Context, monitorId uuid.UUID, from, to time.Time) ([]*CheckResult, *errs.AppError)
	GetLastMonitorCheckResult(ctx context.Context, monitorId uuid.UUID) (*CheckResult, *errs.AppError)
	CreateCheckResult(context.Context, *CheckResult) *errs.AppError
	DeleteCheckResult(context.Context, uuid.UUID) *errs.AppError
}

type RefreshTokenRepo interface {
	CreateToken(context.Context, *RefreshToken) *errs.AppError
	GetToken(context.Context, string) (*RefreshToken, *errs.AppError)
	RevokeToken(context.Context) *errs.AppError
	DeleteToken(context.Context) *errs.AppError
	DeleteUserTokens(context.Context, uuid.UUID) *errs.AppError
}
