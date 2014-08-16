package models

import (
	"time"

	"github.com/elos/server/db"
	"gopkg.in/mgo.v2/bson"
)

const EventKind db.Kind = "event"

type Event struct {
	// Core
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`

	// Properties
	Name      string    `json:"name"`
	StartTime time.Time `json:"start_time" bson:"start_time"`
	EndTime   time.Time `json:"end_time" bson:"end_time"`

	// Links
	User   Link          `json:"user" bson:"user"`
	UserId bson.ObjectId `json:"user_id" bson:"user_id"`
}

func (e *Event) GetId() bson.ObjectId {
	return e.Id
}

func (e *Event) SetId(id bson.ObjectId) {
	e.Id = id
}

func (e *Event) Kind() db.Kind {
	return EventKind
}

func (e *Event) Save() error {
	e.SyncRelationships()

	err := db.Save(e)

	if err == nil {
		e.DidSave()
	}

	return err
}

// Manages the relationship on the other models
func (e *Event) SyncRelationships() error {
	user, err := FindUser(e.UserId)

	if err != nil {
		return err
	}

	return user.AddEvent(e)
}

func (e *Event) Concerned() []bson.ObjectId {
	a := make([]bson.ObjectId, 1)
	a[0] = e.UserId
	return a
}

func (e *Event) GetUser() *User {
	user := User{}

	if e.UserId == "" {
		return &user
	}

	user.Id = e.UserId

	db.PopulateById(&user)

	return &user
}

func (e *Event) SetUser(user *User) error {
	e.UserId = user.GetId()
	return e.Save()
}
