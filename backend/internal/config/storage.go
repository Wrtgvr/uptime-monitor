package config

// Database config is about connection to DB
// Storage config is about storage rules

type StorageConfig struct {
	MaxMonitorsPerUser   int `yaml:"max_monitors_per_user"`
	MinCheckInterval     int `yaml:"min_check_interval"` // in seconds
	DefaultRetentionDays int `yaml:"default_retention_days"`
}
