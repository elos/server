package db

import "gopkg.in/mgo.v2/bson"

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
	Link(string, Model) error

	// Persistence
	Save() error

	// For model updates
	Concerned() []bson.ObjectId
}
