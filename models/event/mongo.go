package event

import (
	"github.com/elos/data"
	"github.com/elos/data/mongo"
	"github.com/elos/server/models"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type MongoEvent struct {
	EID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	ECreatedAt time.Time     `json:"created_at" bson:"created_at"`
	EUpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
	EName      string        `json:"name" bson:"name"`
	EStartTime time.Time     `json:"start_time" bson:"start_time"`
	EEndTime   time.Time     `json:"end_time" bson:"end_time"`
	UserID     bson.ObjectId `json:"user_id" bson:"user_id,omitempty"`
}

func (e *MongoEvent) Kind() data.Kind {
	return CurrentEventKind
}

func (e *MongoEvent) Schema() data.Schema {
	return CurrentEventSchema
}

func (e *MongoEvent) Version() int {
	return CurrentEventVersion
}

func (e *MongoEvent) Valid() bool {
	valid, _ := Validate(e)
	return valid
}

func (u *MongoEvent) DBType() data.DBType {
	return mongo.DBType
}

func (e *MongoEvent) Concerned() []data.ID {
	a := make([]data.ID, 1)
	a[0] = e.UserID
	return a
}

func (e *MongoEvent) SetUser(u models.User) error {
	return e.Schema().Link(e, u)
}

func (e *MongoEvent) LinkOne(r data.Model) {
	switch r.(type) {
	case models.User:
		e.UserID = r.ID().(bson.ObjectId)
	default:
		return
	}
}

func (e *MongoEvent) LinkMul(r data.Model) {
	return
}

func (e *MongoEvent) UnlinkOne(r data.Model) {
	switch r.(type) {
	case models.User:
		if e.UserID == r.ID() {
			e.UserID = *new(bson.ObjectId)
		}
	default:
		return
	}
}

func (e *MongoEvent) UnlinkMul(r data.Model) {
	return
}

// Accessors

func (e *MongoEvent) ID() data.ID {
	return e.EID
}

func (e *MongoEvent) SetID(id data.ID) {
	vid, ok := id.(bson.ObjectId)
	if ok {
		e.EID = vid
	}
}

func (e *MongoEvent) CreatedAt() time.Time {
	return e.ECreatedAt
}

func (e *MongoEvent) SetCreatedAt(t time.Time) {
	e.ECreatedAt = t
}

func (e *MongoEvent) UpdatedAt() time.Time {
	return e.EUpdatedAt
}

func (e *MongoEvent) SetUpdatedAt(t time.Time) {
	e.EUpdatedAt = t
}

func (e *MongoEvent) Name() string {
	return e.EName
}

func (e *MongoEvent) SetName(n string) {
	e.EName = n
}

func (e *MongoEvent) StartTime() time.Time {
	return e.EStartTime
}

func (e *MongoEvent) SetStartTime(t time.Time) {
	e.EStartTime = t
}

func (e *MongoEvent) EndTime() time.Time {
	return e.EEndTime
}

func (e *MongoEvent) SetEndTime(t time.Time) {
	e.EEndTime = t
}
