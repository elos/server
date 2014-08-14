package db

import "gopkg.in/mgo.v2/bson"

type Kind string

type Key string

type Property interface {
}

type Model interface {
	// Basic
	SetId(bson.ObjectId)
	GetId() bson.ObjectId

	// Persistence
	Save() error

	// For model updates
	Concerned() []bson.ObjectId
}
