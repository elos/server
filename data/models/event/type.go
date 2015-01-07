package event

import (
	"errors"
	"time"

	"github.com/elos/server/data"
	"gopkg.in/mgo.v2/bson"
)

// Definition {{{

const Kind data.Kind = "event"

type Event struct {
	// Core
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`

	// Properties
	Name      string    `json:"name"`
	StartTime time.Time `json:"start_time" bson:"start_time"`
	EndTime   time.Time `json:"end_time" bson:"end_time"`

	// Links
	// User   Link          `json:"user" bson:"user"`
	UserId bson.ObjectId `json:"user_id" bson:"user_id,omitempty"`
}

// }}}

// Constructors {{{

func New() *Event {
	return &Event{}
}

func Create(name string, userIdString string) (data.Model, error) {
	if !bson.IsObjectIdHex(userIdString) {
		return nil, errors.New("Invalid userId")
	}

	userId := bson.ObjectIdHex(userIdString)
	if err := data.CheckId(userId); err != nil {
		return nil, err
	}

	event := &Event{
		Id:        userId,
		CreatedAt: time.Now(),
		Name:      name,
	}

	/*
		user, _ := models.Find(models.UserKind, bson.ObjectIdHex(userId))

		db.PopulateById(user)

		event.Link("user", user)
	*/

	if err := event.Save(); err != nil {
		return nil, err
	} else {
		return event, nil
	}
}

// }}}

// Type Methods {{{

func Find(id bson.ObjectId) (data.Model, error) {

	event := New()
	event.Id = id

	if err := db.PopulateById(event); err != nil {
		return event, err
	}

	return event, nil
}

func FindEventBy(field string, value interface{}) (data.Model, error) {
	event := &Event{}

	if err := db.PopulateByField(field, value, event); err != nil {
		return event, err
	}

	return event, nil
}

// }}}
