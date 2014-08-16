package event

import (
	"time"

	"github.com/elos/server/db"
	"github.com/elos/server/models"
	"gopkg.in/mgo.v2/bson"
)

const Kind db.Kind = models.EventKind

func New() *models.Event {
	return &models.Event{}
}

func Create(name string, userId string) (db.Model, error) {
	event := &models.Event{
		Id:        bson.ObjectIdHex(userId),
		CreatedAt: time.Now(),
		Name:      name,
	}

	user, _ := models.FindUser(bson.ObjectIdHex(userId))

	db.PopulateById(user)

	event.SetUser(user)

	if err := event.Save(); err != nil {
		return nil, err
	} else {
		return event, nil
	}
}

func Find(id bson.ObjectId) (db.Model, error) {

	event := New()
	event.Id = id

	if err := db.PopulateById(event); err != nil {
		return event, err
	}

	return event, nil
}
