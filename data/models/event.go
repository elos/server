package models

type Event interface {
	Model
	Nameable
	Timeable

	SetUser(User) error
}
