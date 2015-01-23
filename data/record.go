package data

import (
	"errors"
)

/*
	Model type or class
	Should correspond with the model name, generally lowercase
*/
type Kind string

type Persistable interface {
	Kind() Kind
	GetID() ID

	// For model updates
	Concerned() []ID
}

type Record interface {
	Persistable

	SetID(ID)
	Save() error
}

func CheckID(id ID) error {
	if !id.Valid() {
		return errors.New("Invalid Id")
	}

	return nil
}
