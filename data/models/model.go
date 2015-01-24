package models

import (
	"github.com/elos/server/data"
	"time"
)

type Versioned interface {
	GetVersion() int
}

type Loaded interface {
	GetLoadedAt() time.Time
}

type Createable interface {
	SetCreatedAt(time.Time)
	GetCreatedAt() time.Time
}

type Updateable interface {
	SetUpdatedAt(time.Time)
	GetUpdatedAt() time.Time
}

type Linkable interface {
	LinkOne(Model)
	LinkMul(Model)
	UnlinkOne(Model)
	UnlinkMul(Model)
}

type Model interface {
	data.Record
	Versioned
	Loaded

	Linkable
	Createable
	Updateable
}

type Nameable interface {
	SetName(string)
	GetName() string
}
