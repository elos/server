package event

import (
	"errors"
	"time"

	"github.com/elos/data"
	"github.com/elos/data/mongo"
	"github.com/elos/server/models"
	"gopkg.in/mgo.v2/bson"
)

func New() models.Event {
	return &MongoEvent{
		Based: &models.Based{},
		Named: &models.Named{},
		Timed: &models.Timed{},
	}
}

func Create(db data.DB, name string, userIdString string) (models.Event, error) {
	if !mongo.IsObjectIDHex(userIdString) {
		return nil, errors.New("Invalid userId")
	}

	userId := mongo.NewObjectIDFromHex(userIdString)

	event := New()

	event.SetID(userId.(bson.ObjectId))
	event.SetCreatedAt(time.Now())
	event.SetName(name)

	// user, _ := db.Find(models.UserKind, data.IDHex(userId))
	//event.Schema().Link(event, user)

	if err := db.Save(event); err != nil {
		return nil, err
	} else {
		return event, nil
	}
}

func Find(db data.DB, id data.ID) (models.Event, error) {
	event := New()
	event.SetID(id.(bson.ObjectId))

	if err := db.PopulateByID(event); err != nil {
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
