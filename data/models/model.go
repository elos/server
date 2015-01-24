package models

import (
	"github.com/elos/server/data"
	"time"
)

type Linker interface {
	Link(Model, Model) error
}

type Validateable interface {
	Valid() bool
}

type Versioned interface {
	GetVersion() int
}

type Schema interface {
	Linker
	Validateable
	Versioned
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

	Schema() Schema
}

type Nameable interface {
	SetName(string)
	GetName() string
}

type Timeable interface {
	SetStartTime(time.Time)
	GetStartTime() time.Time
	SetEndTime() time.Time
	GetEndTime() time.Time
}
