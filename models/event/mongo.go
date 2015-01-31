package event

import (
	"github.com/elos/data"
	"github.com/elos/data/mongo"
	"github.com/elos/server/models"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type mongoEvent struct {
	EID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	ECreatedAt time.Time     `json:"created_at" bson:"created_at"`
	EUpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
	EName      string        `json:"name" bson:"name"`
	EStartTime time.Time     `json:"start_time" bson:"start_time"`
	EEndTime   time.Time     `json:"end_time" bson:"end_time"`
	UserID     bson.ObjectId `json:"user_id" bson:"user_id,omitempty"`
}

func (e *mongoEvent) Kind() data.Kind {
	return kind
}

func (e *mongoEvent) Schema() data.Schema {
	return schema
}

func (e *mongoEvent) Version() int {
	return version
}

func (e *mongoEvent) Valid() bool {
	valid, _ := Validate(e)
	return valid
}

func (u *mongoEvent) DBType() data.DBType {
	return mongo.DBType
}

func (e *mongoEvent) Concerned() []data.ID {
	a := make([]data.ID, 1)
	a[0] = e.UserID
	return a
}

func (e *mongoEvent) SetUser(u models.User) error {
	return e.Schema().Link(e, u)
}

func (e *mongoEvent) LinkOne(r data.Model) {
	switch r.(type) {
	case models.User:
		e.UserID = r.ID().(bson.ObjectId)
	default:
		return
	}
}

func (e *mongoEvent) LinkMul(r data.Model) {
	return
}

func (e *mongoEvent) UnlinkOne(r data.Model) {
	switch r.(type) {
	case models.User:
		if e.UserID == r.ID() {
			e.UserID = *new(bson.ObjectId)
		}
	default:
		return
	}
}

func (e *mongoEvent) UnlinkMul(r data.Model) {
	return
}

// Accessors

func (e *mongoEvent) ID() data.ID {
	return e.EID
}

func (e *mongoEvent) SetID(id data.ID) {
	vid, ok := id.(bson.ObjectId)
	if ok {
		e.EID = vid
	}
}

func (e *mongoEvent) CreatedAt() time.Time {
	return e.ECreatedAt
}

func (e *mongoEvent) SetCreatedAt(t time.Time) {
	e.ECreatedAt = t
}

func (e *mongoEvent) UpdatedAt() time.Time {
	return e.EUpdatedAt
}

func (e *mongoEvent) SetUpdatedAt(t time.Time) {
	e.EUpdatedAt = t
}

func (e *mongoEvent) Name() string {
	return e.EName
}

func (e *mongoEvent) SetName(n string) {
	e.EName = n
}

func (e *mongoEvent) StartTime() time.Time {
	return e.EStartTime
}

func (e *mongoEvent) SetStartTime(t time.Time) {
	e.EStartTime = t
}

func (e *mongoEvent) EndTime() time.Time {
	return e.EEndTime
}

func (e *mongoEvent) SetEndTime(t time.Time) {
	e.EEndTime = t
}
