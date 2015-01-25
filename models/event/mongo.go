package event

import (
	"time"

	"github.com/elos/server/data"
	"github.com/elos/server/models"
	"github.com/elos/server/schema"
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
	UserID bson.ObjectId `json:"user_id" bson:"user_id,omitempty"`
}

func (e *MongoEvent) Save(db data.DB) error {
	valid, err := Validate(e)

	if valid {
		return db.Save(e)
	}

	return err
}

func (e *MongoEvent) Concerned() []data.ID {
	a := make([]data.ID, 1)
	a[0] = e.UserID
	return a
}

func (e *MongoEvent) SetUser(u models.User) error {
	return e.Schema().Link(e, u)
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
	return CurrentEventKind
}

func (e *MongoEvent) LinkOne(r schema.Model) {
	switch r.(type) {
	case models.User:
		e.UserID = r.GetID().(bson.ObjectId)
	default:
		return
	}
}

func (e *MongoEvent) LinkMul(r schema.Model) {
	return
}

func (e *MongoEvent) UnlinkMul(r schema.Model) {
	return
}

func (e *MongoEvent) UnlinkOne(r schema.Model) {
	switch r.(type) {
	case models.User:
		if e.UserID == r.GetID() {
			e.UserID = ""
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

func (e *MongoEvent) Schema() schema.Schema {
	return CurrentEventSchema
}

func (e *MongoEvent) GetEndTime() time.Time {
	return e.EndTime
}

func (e *MongoEvent) SetEndTime(t time.Time) {
	e.EndTime = t
}

func (e *MongoEvent) SetStartTime(t time.Time) {
	e.StartTime = t
}

func (e *MongoEvent) GetStartTime() time.Time {
	return e.StartTime
}

func (e *MongoEvent) SetName(n string) {
	e.Name = n
}

func (e *MongoEvent) GetName() string {
	return e.Name
}

func (e *MongoEvent) Valid() bool {
	return true
}
