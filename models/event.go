package models

import (
	"time"

	"github.com/elos/server/db"
	"github.com/elos/server/hub"
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
	UserId bson.ObjectId `json:"user_id" bson:"user_id"`
}

func (e *Event) GetId() *bson.ObjectId {
	return &e.Id
}

func (e *Event) Save() error {
	err := db.Save(EventKind, e)

	if err != nil {
		e.DidSave()
	}

	return err
}

func (e *Event) DidSave() {
	hub.PrimaryHermes.Send <- e
}

func (e *Event) Concerned() []*bson.ObjectId {
	a := make([]*bson.ObjectId, 1)
	a[0] = &e.UserId
	return a
}

func (e *Event) GetUser() *User {
	user := User{}

	if e.UserId == "" {
		return &user
	}

	user.Id = e.UserId

	db.PopulateById(UserKind, &user)

	return &user
}

func (e *Event) SetUser(user *User) error {
	e.UserId = *user.GetId()

	if !user.EventIdsHash()[&e.Id] {
		user.AddEvent(e)
	}

	return e.Save()
}

func CreateEvent(name string /*startTime time.Time, endTime time.Time,*/, userId string) (*Event, error) {
	event := Event{
		Id:        bson.NewObjectId(),
		CreatedAt: time.Now(),
		Name:      name,
		/*StartTime: startTime,
		EndTime:   endTime,*/
	}

	event.SetUser(&User{Id: bson.ObjectIdHex(userId)})

	if err := event.Save(); err != nil {
		return nil, err
	} else {
		return &event, nil
	}
}
