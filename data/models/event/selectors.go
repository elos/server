package event

import (
	"github.com/elos/server/data"
	"gopkg.in/mgo.v2/bson"
)

func (e *Event) GetID() data.ID {
	return e.ID
}

func (e *Event) SetID(id data.ID) {
	vid, ok := id.(bson.ObjectId)
	if ok {
		e.ID = vid
	}
}

func (e *Event) Kind() data.Kind {
	return Kind
}
