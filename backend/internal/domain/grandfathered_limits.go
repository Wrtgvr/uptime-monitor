package domain

import "time"

// keep limits some time after changing limits config
type GrandfatheredLimits struct {
	MaxMonitors, RetentionDays int
	Reason                     string
	ExpiresAt                  time.Time
}
