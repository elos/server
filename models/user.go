package models

import (
	"github.com/elos/data"
)

type User interface {
	data.Model
	data.Nameable

	SetKey(string)
	Key() string

	AddEvent(Event) error
	RemoveEvent(Event) error
}
