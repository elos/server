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

func (e *Event) Link(property string, model db.Model) {
	switch property {
	case "user":
		id := model.GetId()

		if e.UserId == id {
			return
		}

		e.UserId = model.GetId()
		e.Save()
	default:
		return
	}
}

func (e *Event) GetLink(property string, model db.Model) {
	switch property {
	case "user":
		model.SetId(e.Id)
		db.PopulateById(model)
	default:
		return
	}
}
