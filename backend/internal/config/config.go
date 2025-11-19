package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Port            int `yaml:"server_port"`
	ShutdownTimeout int `yaml:"server_shutdown_timeout"`
}

type Config struct {
	Env      string
	Server   *ServerConfig
	Database *DatabaseConfig
	Redis    *RedisConfig
}

func LoadConfig() *Config {
	//* get app env
	env := os.Getenv(APP_ENV)
	if env == "" {
		log.Fatalf("env variable %s is missing", APP_ENV)
	}
	switch env {
	case AppEnvLocal, AppEnvDev, AppEnvStaging, AppEnvProd:
	default:
		log.Printf("WARN: unexpected app environment: app_env=%s", env)
	}

	//* get config file
	cfgFile, err := os.ReadFile(GetConfigPath(env))
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	//* parse config
	var cfg Config

	if err = yaml.Unmarshal(cfgFile, &cfg); err != nil {
		log.Fatalf("Error unmarshaling config file: %v", err)
	}

	//* load env variables into config struct
	loadRedisEnvVarsIntoConfig(cfg.Redis)
	loadDatabaseEnvVarsIntoConfig(cfg.Database)

	cfg.Env = env

	return &cfg
}

func GetConfigPath(env string) string {
	filename := fmt.Sprintf("config_%s.yaml", env)

	//* 1. Check env variable
	if envPath := os.Getenv(CONFIG_PATH); envPath != "" {
		if exists(envPath) {
			return envPath
		}
		log.Printf("WARN: %s is set but file doesn't exists: %s", CONFIG_PATH, envPath)
	}

	var checkedPaths []string

	//* 2. Near executable/bin
	if exePath, err := os.Executable(); err == nil {
		exeDir := filepath.Dir(exePath)
		configPath := filepath.Join(exeDir, "configs", filename)

		checkedPaths = append(checkedPaths, configPath)

		if exists(configPath) {
			return configPath
		}
	}

	//* 3. Check in work dir
	if cwd, err := os.Getwd(); err == nil {
		configPath := filepath.Join(cwd, "configs", filename)

		checkedPaths = append(checkedPaths, configPath)

		if exists(configPath) {
			return configPath
		}
	}

	log.Printf("Config file '%s' not found in:\n%s", filename, strings.Join(checkedPaths, "\n"))
	log.Fatalf("Please set config path environment variable (%s) or place config in one of the above locations", CONFIG_PATH)

	return ""
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
