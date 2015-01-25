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
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`

	Name     string      `json:"name"`
	Key      string      `json:"key"`
	EventIDs mongo.IDSet `json:"event_ids" bson:"event_ids"`
}

func (u *MongoUser) Kind() data.Kind {
	return CurrentUserKind
}

func (u *MongoUser) SetID(id data.ID) {
	vid, ok := id.(bson.ObjectId)
	if ok {
		u.ID = vid
	}
}

func (u *MongoUser) GetID() data.ID {
	return u.ID
}

func (u *MongoUser) SetName(name string) {
	u.Name = name
}

func (u *MongoUser) GetName() string {
	return u.Name
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
	a[0] = u.ID
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
	u.CreatedAt = t
}

func (u *MongoUser) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *MongoUser) SetUpdatedAt(t time.Time) {
	u.UpdatedAt = t
}

func (u *MongoUser) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

func (u *MongoUser) SetKey(s string) {
	u.Key = s
}

func (u *MongoUser) GetKey() string {
	return u.Key
}

func (u *MongoUser) LinkOne(r schema.Model) {
	return
}

func (u *MongoUser) LinkMul(r schema.Model) {
	switch r.(type) {
	case models.Event:
		u.LinkEvent(r.GetID().(bson.ObjectId))
	default:
		return
	}
}

func (u *MongoUser) UnlinkMul(r schema.Model) {
	switch r.(type) {
	case models.Event:
		u.UnlinkEvent(r.GetID().(bson.ObjectId))
	default:
		return
	}
}

func (u *MongoUser) UnlinkOne(r schema.Model) {
	return
}

func (u *MongoUser) GetVersion() int {
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
