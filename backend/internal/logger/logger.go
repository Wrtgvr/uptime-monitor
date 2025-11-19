package logger

import (
	"log/slog"
	"os"

	"github.com/wrtgvr/uptime-monitor/internal/config"
)

func NewLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case config.AppEnvLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case config.AppEnvDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case config.AppEnvStaging, config.AppEnvProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	default:
		log = NewDefaultLogger()
		log.Warn("unexpected app environment")
		log.Warn("WARN: using default logger", "env", env)
	}

	return log
}

func NewDefaultLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(
		os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		},
	))
}
