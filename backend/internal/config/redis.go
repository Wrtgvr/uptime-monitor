package config

import "os"

type RedisConfig struct {
	Host     string `yaml:"redis_host"`
	Port     int    `yaml:"redis_port"`
	DBNumber int    `yaml:"redis_db_number"`
	Password string
}

func loadRedisEnvVarsIntoConfig(cfg *RedisConfig) {
	cfg.Password = os.Getenv(REDIS_PASSWORD)
}
