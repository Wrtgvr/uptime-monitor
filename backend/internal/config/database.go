package config

// Database config is about connection to DB
// Storage config is about storage rules

import "os"

type DatabaseConfig struct {
	Host     string `yaml:"db_host"`
	Port     int    `yaml:"db_port"`
	User     string
	Password string
	Name     string
}

func loadDatabaseEnvVarsIntoConfig(cfg *DatabaseConfig) {
	cfg.User = os.Getenv(DB_USER)
	cfg.Password = os.Getenv(DB_PASSWORD)
	cfg.Name = os.Getenv(DB_NAME)
}
