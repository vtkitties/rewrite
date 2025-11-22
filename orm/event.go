package orm

import (
	"time"

	"gorm.io/gorm"
)

type Status string

const (
	EventStatusScheduled Status = "scheduled"
	EventStatusOngoing   Status = "ongoing"
	EventStatusCompleted Status = "completed"
)

type Event struct {
	gorm.Model
	Title       string
	Description string
	StartTime   time.Time
	EndTime     time.Time
	Status      Status
}
