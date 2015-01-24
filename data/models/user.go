package models

type User interface {
	Model
	Nameable

	SetKey(string)
	GetKey() string

	AddEvent(Event) error
	RemoveEvent(Event) error
}
