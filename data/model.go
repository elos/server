package data

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

/*
	Model type or class
	Should correspond with the model name, generally lowercase
*/
type Kind string

type Model interface {
	// Core
	SetId(bson.ObjectId)
	GetId() bson.ObjectId
	Kind() Kind

	// Persistence
	Save() error

	// For model updates
	Concerned() []bson.ObjectId
}

func CheckId(id bson.ObjectId) error {
	if !id.Valid() {
		return errors.New("Invalid Id")
	}

	return nil
}
