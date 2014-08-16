package models

import (
	"fmt"

	"github.com/elos/server/db"
	"gopkg.in/mgo.v2/bson"
)

type Link struct {
	Id   bson.ObjectId `json:"id" bson:"id"`
	Kind db.Kind       `json:"=" bson:"-"`
}

func (l *Link) Populate() (db.Model, error) {
	var model db.Model

	switch l.Kind {
	case UserKind:
		model = &User{Id: l.Id}
	case EventKind:
		model = &Event{Id: l.Id}
	default:
		return &User{Id: l.Id}, fmt.Errorf("Kind not recognized")
	}

	return model, nil
}
