package main

import (
	"github.com/wrtgvr/uptime-monitor/internal/config"
	"github.com/wrtgvr/uptime-monitor/internal/logger"
)

func main() {
	cfg := config.LoadConfig()
	log := logger.NewLogger(cfg.Env)

	log.Info("Successfully got config", "server_port", cfg.Server.Port)
}
