package db

import "gopkg.in/mgo.v2/bson"

/*
	Model type or class
	 - Mongo: Collection
	 - Relational: Row
	 Should correspond with the type name, generally plural lowercase
*/
type Kind string

// Not used
type Key string

type Model interface {
	// Core
	SetId(bson.ObjectId)
	GetId() bson.ObjectId
	Kind() Kind

	// Persistence
	Save() error
	Link(string, Model)

	// For model updates
	Concerned() []bson.ObjectId
}

type Link struct {
	Id   bson.ObjectId `json:"id" bson:"id"`
	Kind Kind          `json:"=" bson:"-"`
}
