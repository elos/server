package event

import (
	"github.com/elos/server/db"
	"gopkg.in/mgo.v2/bson"
)

func (e *Event) Save() error {
	e.SyncRelationships()

	err := db.Save(e)

	if err == nil {
		// e.DidSave()
	}

	return err
}

// Manages the relationship on the other models
func (e *Event) SyncRelationships() error {
	/*
		model, err := models.FindUser(e.UserId)

		if err != nil {
			return err
		}

		model.Link("event", e)
	*/

	// TODO: fix
	return nil
}

func (e *Event) Concerned() []bson.ObjectId {
	a := make([]bson.ObjectId, 1)
	a[0] = e.UserId
	return a
}

func (e *Event) Link(property string, model db.Model) {
	switch property {
	case "user":
		if e.UserId == model.GetId() {
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
