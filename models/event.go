package models

import (
	"github.com/elos/data"
)

type Event interface {
	data.Model
	data.Nameable
	data.Timeable

	SetUser(User) error
}
