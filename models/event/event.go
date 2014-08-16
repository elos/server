package event

import (
	"time"

	"github.com/elos/server/db"
	"gopkg.in/mgo.v2/bson"
)

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
	UserId bson.ObjectId `json:"user_id" bson:"user_id"`
}

func (e *Event) GetId() bson.ObjectId {
	return e.Id
}

func (e *Event) SetId(id bson.ObjectId) {
	e.Id = id
}

func (e *Event) Kind() db.Kind {
	return Kind
}

func (e *Event) Save() error {
	e.SyncRelationships()

	err := db.Save(e)

	if err == nil {
		// e.DidSave()
	}

	return err
}

// Manages the relationship on the other models
func (e *Event) SyncRelationships() error {
	/*
		model, err := models.FindUser(e.UserId)

		if err != nil {
			return err
		}

		model.Link("event", e)
	*/

	// TODO: fix
	return nil
}

func (e *Event) Concerned() []bson.ObjectId {
	a := make([]bson.ObjectId, 1)
	a[0] = e.UserId
	return a
}

func (e *Event) Link(property string, model db.Model) {
	switch property {
	case "user":
		if e.UserId == model.GetId() {
			return
		}

		e.UserId = model.GetId()
		e.Save()
	default:
		return
	}
}

func (e *Event) GetLink(property string, model db.Model) {
	switch property {
	case "user":
		model.SetId(e.Id)
		db.PopulateById(model)
	default:
		return
	}
}

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

func Find(id bson.ObjectId) (db.Model, error) {

	event := New()
	event.Id = id

	if err := db.PopulateById(event); err != nil {
		return event, err
	}

	return event, nil
}
