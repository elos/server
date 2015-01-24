package event

import (
	"time"

	"github.com/elos/server/data"
	"github.com/elos/server/data/models"
	"gopkg.in/mgo.v2/bson"
)

type MongoEvent struct {
	// Core
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
	LoadedAt  time.Time     `json:"-" bson:"-"`

	// Properties
	Name      string    `json:"name"`
	StartTime time.Time `json:"start_time" bson:"start_time"`
	EndTime   time.Time `json:"end_time" bson:"end_time"`

	// Links
	// User   Link          `json:"user" bson:"user"`
	UserID data.ID `json:"user_id" bson:"user_id,omitempty"`
}

func (e *MongoEvent) Save() error {
	return db.Save(e)
}

func (e *MongoEvent) Concerned() []data.ID {
	a := make([]data.ID, 1)
	a[0] = e.UserID
	return a
}

func (e *MongoEvent) SetUser(userId data.ID) error {
	if err := data.CheckID(userId); err != nil {
		return err
	}

	if e.UserID == userId {
		return nil
	}

	e.UserID = userId

	return e.Save()
}

func (e *MongoEvent) GetID() data.ID {
	return e.ID
}

func (e *MongoEvent) SetID(id data.ID) {
	vid, ok := id.(bson.ObjectId)
	if ok {
		e.ID = vid
	}
}

func (e *MongoEvent) Kind() data.Kind {
	return models.EventKind
}

func (e *MongoEvent) LinkOne(r models.Model) {
	switch r.Kind() {
	case models.UserKind:
		e.UserID = r.GetID()
		e.Save()
	default:
		return
	}
}

func (e *MongoEvent) LinkMul(r models.Model) {
	return
}

func (e *MongoEvent) UnlinkMul(r models.Model) {
	return
}

func (e *MongoEvent) UnlinkOne(r models.Model) {
	switch r.Kind() {
	case models.UserKind:
		if e.UserID == r.GetID() {
			e.UserID = nil
			e.Save()
		}
	default:
		return
	}
}

func (u *MongoEvent) GetLoadedAt() time.Time {
	return u.LoadedAt
}

func (e *MongoEvent) GetCreatedAt() time.Time {
	return e.CreatedAt
}

func (e *MongoEvent) SetCreatedAt(t time.Time) {
	e.CreatedAt = t
}

func (e *MongoEvent) GetUpdatedAt() time.Time {
	return e.UpdatedAt
}

func (e *MongoEvent) SetUpdatedAt(t time.Time) {
	e.UpdatedAt = t
}

func (e *MongoEvent) GetVersion() int {
	return CurrentEventVersion
}

func (e *MongoEvent) Schema() models.Schema {
	return CurrentEventSchema
}
