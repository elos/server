package models

import (
	"time"
)

type Event interface {
	Model
	Nameable

	SetStartTime(time.Time)
	GetStartTime() time.Time
	SetEndTime() time.Time
	GetEndTime() time.Time
}
