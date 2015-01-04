package data

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

/*
	Model type or class
	 - Mongo: Collection
	 - Relational: Row
	 Should correspond with the type name, generally plural lowercase
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
		return fmt.Errorf("Invalid Id")
	}

	return nil
}
