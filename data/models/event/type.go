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
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`

	// Properties
	Name      string    `json:"name"`
	StartTime time.Time `json:"start_time" bson:"start_time"`
	EndTime   time.Time `json:"end_time" bson:"end_time"`

	// Links
	// User   Link          `json:"user" bson:"user"`
	UserId data.ID `json:"user_id" bson:"user_id,omitempty"`
}

// }}}

// Constructors {{{

func New() *Event {
	return &Event{}
}

func Create(name string, userIdString string) (data.Model, error) {
	if !data.IsObjectIDHex(userIdString) {
		return nil, errors.New("Invalid userId")
	}

	userId := data.NewObjectIDFromHex(userIdString)
	if err := data.CheckID(userId); err != nil {
		return nil, err
	}

	event := &Event{
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

// }}}

// Type Methods {{{

func Find(id data.ID) (data.Model, error) {
	event := New()
	event.ID = id.(bson.ObjectId)

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
