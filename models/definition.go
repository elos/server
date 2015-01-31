/*
	The Models Package supplies the domain specific data interfaces elos relies on.

	It's subdirectories supply the implementation for the interfaces defined here.
*/
package models

import "github.com/elos/data"

type User interface {
	data.Model
	data.Nameable

	SetKey(string)
	Key() string

	AddEvent(Event) error
	RemoveEvent(Event) error
}

type Event interface {
	data.Model
	data.Nameable
	data.Timeable

	SetUser(User) error
}
