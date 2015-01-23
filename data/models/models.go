package models

import (
	"github.com/elos/server/data"
	"time"
)

const UserKind data.Kind = "user"
const EventKind data.Kind = "event"

type Model interface {
	SetCreatedAt(time.Time)
	GetCreatedAt() time.Time
	SetUpdatedAt(time.Time)
	GetUpdatedAt() time.Time

	GetLoadedAt() time.Time
}

type Nameable interface {
	SetName(string)
	GetName() string
}

type User interface {
	data.Record
	Model
	Nameable

	SetKey(string)
	GetKey() string
}

type Event interface {
	data.Record
	Model

	Nameable

	SetStartTime(time.Time)
	GetStartTime() time.Time
	SetEndTime() time.Time
	GetEndTime() time.Time
}
