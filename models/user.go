package models

import (
	"github.com/elos/server/schema"
)

type User interface {
	schema.Model
	schema.Nameable

	SetKey(string)
	GetKey() string

	AddEvent(Event) error
	RemoveEvent(Event) error
}
