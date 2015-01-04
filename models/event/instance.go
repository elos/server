package event

import (
	"github.com/elos/server/db"
	"gopkg.in/mgo.v2/bson"
)

func (e *Event) Save() error {
	return db.Save(e)
}

func (e *Event) Concerned() []bson.ObjectId {
	a := make([]bson.ObjectId, 1)
	a[0] = e.UserId
	return a
}

func (e *Event) SetUser(userId bson.ObjectId) error {
	if err := db.CheckId(userId); err != nil {
		return err
	}

	if e.UserId == userId {
		return nil
	}

	e.UserId = userId

	return e.Save()
}
