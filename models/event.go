package models

import (
	"github.com/elos/server/schema"
)

type Event interface {
	schema.Model
	schema.Nameable
	schema.Timeable

	SetUser(User) error
}
