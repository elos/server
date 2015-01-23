package event

import (
	"errors"
	"time"

	"github.com/elos/server/data"
	"github.com/elos/server/data/mongo"
	"gopkg.in/mgo.v2/bson"
)

func New() data.Record {
	return &MongoEvent{}
}

func Create(name string, userIdString string) (data.Record, error) {
	if !mongo.IsObjectIDHex(userIdString) {
		return nil, errors.New("Invalid userId")
	}

	userId := mongo.NewObjectIDFromHex(userIdString)
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

func Find(id data.ID) (data.Record, error) {
	event := New()
	event.SetID(id.(bson.ObjectId))

	if err := db.PopulateById(event); err != nil {
		return event, err
	}

	return event, nil
}

func FindEventBy(field string, value interface{}) (data.Record, error) {
	event := &MongoEvent{}

	if err := db.PopulateByField(field, value, event); err != nil {
		return event, err
	}

	return event, nil
}
