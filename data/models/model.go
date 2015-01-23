package models

import (
	"github.com/elos/server/data"
	"time"
)

type Model interface {
	data.Record

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
