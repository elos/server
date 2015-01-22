package event

import (
	"errors"
	"time"

	"github.com/elos/server/data"
	"gopkg.in/mgo.v2/bson"
)

func New() data.Model {
	return &MongoEvent{}
}

func Create(name string, userIdString string) (data.Model, error) {
	if !data.IsObjectIDHex(userIdString) {
		return nil, errors.New("Invalid userId")
	}

	userId := data.NewObjectIDFromHex(userIdString)
	if err := data.CheckID(userId); err != nil {
		return nil, err
	}

	event := &MongoEvent{
		ID:        userId.(bson.ObjectId),
		CreatedAt: time.Now(),
		Name:      name,
	}

	/*
		user, _ := models.Find(models.UserKind, data.IDHex(userId))

		db.PopulateById(user)

		event.Link("user", user)
	*/

	if err := event.Save(); err != nil {
		return nil, err
	} else {
		return event, nil
	}
}

func Find(id data.ID) (data.Model, error) {
	event := New()
	event.SetID(id.(bson.ObjectId))

	if err := db.PopulateById(event); err != nil {
		return event, err
	}

	return event, nil
}

func FindEventBy(field string, value interface{}) (data.Model, error) {
	event := &MongoEvent{}

	if err := db.PopulateByField(field, value, event); err != nil {
		return event, err
	}

	return event, nil
}
