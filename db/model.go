package db

import "gopkg.in/mgo.v2/bson"

type Kind string

type Key string

type Property interface {
}

type Model interface {
	GetId() *bson.ObjectId
	Save() error
	Concerned() []*bson.ObjectId
}
