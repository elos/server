package event

import (
	"time"

	"github.com/elos/server/db"
	"gopkg.in/mgo.v2/bson"
)

// Definition {{{

var DB db.DB

const Kind db.Kind = "event"

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

func Create(name string, userId string) (db.Model, error) {
	event := &Event{
		Id:        bson.ObjectIdHex(userId),
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

func Find(id bson.ObjectId) (db.Model, error) {

	event := New()
	event.Id = id

	if err := DB.PopulateById(event); err != nil {
		return event, err
	}

	return event, nil
}

func FindEventBy(field string, value interface{}) (db.Model, error) {
	event := &Event{}

	if err := DB.PopulateByField(field, value, event); err != nil {
		return event, err
	}

	return event, nil
}

// }}}
