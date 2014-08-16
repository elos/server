package event

import (
	"time"

	"github.com/elos/server/db"
	"github.com/elos/server/models"
	"gopkg.in/mgo.v2/bson"
)

func New() *models.Event {
	return &models.Event{}
}

func Create(name string, userId string) (*models.Event, error) {
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

func Find(id bson.ObjectId) (*models.Event, error) {

	event := New()
	event.Id = id

	if err := db.PopulateById(event); err != nil {
		return event, err
	}

	return event, nil
}
