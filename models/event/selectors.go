package event

import (
	"github.com/elos/server/db"
	"gopkg.in/mgo.v2/bson"
)

func (e *Event) GetId() bson.ObjectId {
	return e.Id
}

func (e *Event) SetId(id bson.ObjectId) {
	e.Id = id
}

func (e *Event) Kind() db.Kind {
	return Kind
}
