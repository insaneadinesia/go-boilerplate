package models

import (
	"time"
)

type HealthCheck struct {
	CurrentTimestamp time.Time `json:"current_timestamp"`
}
