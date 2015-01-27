package user

import (
	"github.com/elos/data"
	"github.com/elos/data/mongo"
	"github.com/elos/schema"
	"github.com/elos/server/models"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type MongoUser struct {
	id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	createdAt time.Time     `json:"created_at" bson:"created_at"`
	updatedAt time.Time     `json:"updated_at" bson:"updated_at"`

	name     string      `json:"name"`
	key      string      `json:"key"`
	eventIDs mongo.IDSet `json:"event_ids" bson:"event_ids"`
}

func (u *MongoUser) Kind() data.Kind {
	return CurrentUserKind
}

func (u *MongoUser) SetID(id data.ID) {
	vid, ok := id.(bson.ObjectId)
	if ok {
		u.id = vid
	}
}

func (u *MongoUser) ID() data.ID {
	return u.id
}

func (u *MongoUser) SetName(name string) {
	u.name = name
}

func (u *MongoUser) Name() string {
	return u.name
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
	a[0] = u.id
	return a
}

func (u *MongoUser) LinkEvent(eventID bson.ObjectId) error {
	u.eventIDs = mongo.AddID(u.eventIDs, eventID)
	return nil
}

func (u *MongoUser) UnlinkEvent(eventID bson.ObjectId) error {
	u.eventIDs = mongo.DropID(u.eventIDs, eventID)
	return nil
}

func (u *MongoUser) SetCreatedAt(t time.Time) {
	u.createdAt = t
}

func (u *MongoUser) CreatedAt() time.Time {
	return u.createdAt
}

func (u *MongoUser) SetUpdatedAt(t time.Time) {
	u.updatedAt = t
}

func (u *MongoUser) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *MongoUser) SetKey(s string) {
	u.key = s
}

func (u *MongoUser) Key() string {
	return u.key
}

func (u *MongoUser) LinkOne(r schema.Model) {
	return
}

func (u *MongoUser) LinkMul(r schema.Model) {
	switch r.(type) {
	case models.Event:
		u.LinkEvent(r.ID().(bson.ObjectId))
	default:
		return
	}
}

func (u *MongoUser) UnlinkMul(r schema.Model) {
	switch r.(type) {
	case models.Event:
		u.UnlinkEvent(r.ID().(bson.ObjectId))
	default:
		return
	}
}

func (u *MongoUser) UnlinkOne(r schema.Model) {
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

func (u *MongoUser) Schema() schema.Schema {
	return CurrentUserSchema
}

func (u *MongoUser) Valid() bool {
	valid, _ := Validate(u)
	return valid
}

func (u *MongoUser) DBType() data.DBType {
	return mongo.DBType
}