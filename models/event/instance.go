package event

import (
	"fmt"
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

func (e *Event) Link(property string, model db.Model) error {
	switch property {
	case "user":
		id := model.GetId()

		if e.UserId == id {
			return nil
		}

		e.UserId = model.GetId()
		e.Save()
		return nil
	default:
		return fmt.Errorf("Unrecognized Property")
	}
}

func (e *Event) GetLink(property string, model db.Model) error {
	switch property {
	case "user":
		model.SetId(e.Id)
		db.PopulateById(model)
		return nil
	default:
		return fmt.Errorf("Unrecognized property")
	}
}
