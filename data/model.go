package data

import (
	"errors"
)

/*
	Model type or class
	Should correspond with the model name, generally lowercase
*/
type Kind string

type Model interface {
	Agent

	SetID(ID)
	Kind() Kind
	Save() error

	// For model updates
	Concerned() []ID
}

func CheckID(id ID) error {
	if !id.Valid() {
		return errors.New("Invalid Id")
	}

	return nil
}
