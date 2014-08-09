package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Persistable interface {
}

type Model struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
}
