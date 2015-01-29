package user

import (
	"github.com/elos/data"
	"github.com/elos/data/mongo"
	"github.com/elos/server/models"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type MongoUser struct {
	EID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	ECreatedAt time.Time     `json:"created_at" bson:"created_at"`
	EUpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`

	EName    string      `json:"name" bson:"name"`
	EKey     string      `json:"key" bson:"key"`
	EventIDs mongo.IDSet `json:"event_ids" bson:"event_ids"`
}

func (u *MongoUser) Kind() data.Kind {
	return CurrentUserKind
}

func (u *MongoUser) SetID(id data.ID) {
	vid, ok := id.(bson.ObjectId)
	if ok {
		u.EID = vid
	}
}

func (u *MongoUser) ID() data.ID {
	return u.EID
}

func (u *MongoUser) SetName(name string) {
	u.EName = name
}

func (u *MongoUser) Name() string {
	return u.EName
}

func (u *MongoUser) Save(db data.DB) error {
	valid, err := Validate(u)
	if valid {
		return db.Save(u)
	} else {
		return err
	}
}

func (u *MongoUser) Concerned() []data.ID {
	a := make([]data.ID, 1)
	a[0] = u.EID
	return a
}

func (u *MongoUser) LinkEvent(eventID bson.ObjectId) error {
	u.EventIDs = mongo.AddID(u.EventIDs, eventID)
	return nil
}

func (u *MongoUser) UnlinkEvent(eventID bson.ObjectId) error {
	u.EventIDs = mongo.DropID(u.EventIDs, eventID)
	return nil
}

func (u *MongoUser) SetCreatedAt(t time.Time) {
	u.ECreatedAt = t
}

func (u *MongoUser) CreatedAt() time.Time {
	return u.ECreatedAt
}

func (u *MongoUser) SetUpdatedAt(t time.Time) {
	u.EUpdatedAt = t
}

func (u *MongoUser) UpdatedAt() time.Time {
	return u.EUpdatedAt
}

func (u *MongoUser) SetKey(s string) {
	u.EKey = s
}

func (u *MongoUser) Key() string {
	return u.EKey
}

func (u *MongoUser) LinkOne(r data.Model) {
	return
}

func (u *MongoUser) LinkMul(r data.Model) {
	switch r.(type) {
	case models.Event:
		u.LinkEvent(r.ID().(bson.ObjectId))
	default:
		return
	}
}

func (u *MongoUser) UnlinkMul(r data.Model) {
	switch r.(type) {
	case models.Event:
		u.UnlinkEvent(r.ID().(bson.ObjectId))
	default:
		return
	}
}

func (u *MongoUser) UnlinkOne(r data.Model) {
	return
}

func (u *MongoUser) Version() int {
	return CurrentUserVersion
}

func (u *MongoUser) AddEvent(e models.Event) error {
	return u.Schema().Link(u, e)
}

func (u *MongoUser) RemoveEvent(e models.Event) error {
	return u.Schema().Unlink(u, e)
}

func (u *MongoUser) Schema() data.Schema {
	return CurrentUserSchema
}

func (u *MongoUser) Valid() bool {
	valid, _ := Validate(u)
	return valid
}

func (u *MongoUser) DBType() data.DBType {
	return mongo.DBType
}
