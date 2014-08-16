package db

import "gopkg.in/mgo.v2/bson"

type Kind string

type Key string

type Model interface {
	// Basic
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
