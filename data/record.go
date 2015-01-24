package data

import (
	"errors"
)

/*
	Model type or class
	Should correspond with the model name, generally lowercase
*/
type Kind string

type Identifiable interface {
	SetID(ID)
	GetID() ID
}

type Persistable interface {
	Identifiable

	// For saving
	Kind() Kind

	// For model updates
	Concerned() []ID
}

type Record interface {
	Persistable

	Save(DB) error
}

func CheckID(id ID) error {
	if !id.Valid() {
		return errors.New("Invalid Id")
	}

	return nil
}
