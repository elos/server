package models

type User interface {
	Model
	Nameable

	SetKey(string)
	GetKey() string
}
