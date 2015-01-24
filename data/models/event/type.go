package event

import (
	"errors"
	"time"

	"github.com/elos/server/data"
	"github.com/elos/server/data/models"
	"github.com/elos/server/data/mongo"
	"github.com/elos/server/data/schema"
	"gopkg.in/mgo.v2/bson"
)

func New() models.Event {
	return &MongoEvent{}
}

func Create(db data.DB, name string, userIdString string) (models.Event, error) {
	if !mongo.IsObjectIDHex(userIdString) {
		return nil, errors.New("Invalid userId")
	}

	userId := mongo.NewObjectIDFromHex(userIdString)

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

	if err := event.Save(db); err != nil {
		return nil, err
	} else {
		return event, nil
	}
}

func Find(db data.DB, id data.ID) (models.Event, error) {
	event := New()
	event.SetID(id.(bson.ObjectId))

	if err := db.PopulateById(event); err != nil {
		return event, err
	}

	return event, nil
}

func FindEventBy(db data.DB, field string, value interface{}) (models.Event, error) {
	event := &MongoEvent{}

	if err := db.PopulateByField(field, value, event); err != nil {
		return event, err
	}

	return event, nil
}

var CurrentEventSchema schema.Schema
var CurrentEventVersion int

func SetupModel(s schema.Schema, v int) {
	CurrentEventSchema = s
	CurrentEventVersion = v
}
