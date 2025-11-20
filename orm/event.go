package orm

import (
	"time"

	"gorm.io/gorm"
)

type Status string

const (
	Scheduled Status = "scheduled"
	Ongoing   Status = "ongoing"
	Completed Status = "completed"
)

type Event struct {
	gorm.Model
	Title       string
	Description string
	StartTime   time.Time
	EndTime     time.Time
	Status      Status
}
