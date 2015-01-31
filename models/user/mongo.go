package user

import (
	"github.com/elos/data"
	"github.com/elos/data/mongo"
	"github.com/elos/server/models"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type mongoUser struct {
	EID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	ECreatedAt time.Time     `json:"created_at" bson:"created_at"`
	EUpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`

	EName    string      `json:"name" bson:"name"`
	EKey     string      `json:"key" bson:"key"`
	EventIDs mongo.IDSet `json:"event_ids" bson:"event_ids"`
}

func (u *mongoUser) Kind() data.Kind {
	return CurrentUserKind
}

func (u *mongoUser) SetID(id data.ID) {
	vid, ok := id.(bson.ObjectId)
	if ok {
		u.EID = vid
	}
}

func (u *mongoUser) ID() data.ID {
	return u.EID
}

func (u *mongoUser) SetName(name string) {
	u.EName = name
}

func (u *mongoUser) Name() string {
	return u.EName
}

func (u *mongoUser) Save(db data.DB) error {
	valid, err := Validate(u)
	if valid {
		return db.Save(u)
	} else {
		return err
	}
}

func (u *mongoUser) Concerned() []data.ID {
	a := make([]data.ID, 1)
	a[0] = u.EID
	return a
}

func (u *mongoUser) LinkEvent(eventID bson.ObjectId) error {
	u.EventIDs = mongo.AddID(u.EventIDs, eventID)
	return nil
}

func (u *mongoUser) UnlinkEvent(eventID bson.ObjectId) error {
	u.EventIDs = mongo.DropID(u.EventIDs, eventID)
	return nil
}

func (u *mongoUser) SetCreatedAt(t time.Time) {
	u.ECreatedAt = t
}

func (u *mongoUser) CreatedAt() time.Time {
	return u.ECreatedAt
}

func (u *mongoUser) SetUpdatedAt(t time.Time) {
	u.EUpdatedAt = t
}

func (u *mongoUser) UpdatedAt() time.Time {
	return u.EUpdatedAt
}

func (u *mongoUser) SetKey(s string) {
	u.EKey = s
}

func (u *mongoUser) Key() string {
	return u.EKey
}

func (u *mongoUser) LinkOne(r data.Model) {
	return
}

func (u *mongoUser) LinkMul(r data.Model) {
	switch r.(type) {
	case models.Event:
		u.LinkEvent(r.ID().(bson.ObjectId))
	default:
		return
	}
}

func (u *mongoUser) UnlinkMul(r data.Model) {
	switch r.(type) {
	case models.Event:
		u.UnlinkEvent(r.ID().(bson.ObjectId))
	default:
		return
	}
}

func (u *mongoUser) UnlinkOne(r data.Model) {
	return
}

func (u *mongoUser) Version() int {
	return CurrentUserVersion
}

func (u *mongoUser) AddEvent(e models.Event) error {
	return u.Schema().Link(u, e)
}

func (u *mongoUser) RemoveEvent(e models.Event) error {
	return u.Schema().Unlink(u, e)
}

func (u *mongoUser) Schema() data.Schema {
	return CurrentUserSchema
}

func (u *mongoUser) Valid() bool {
	valid, _ := Validate(u)
	return valid
}

func (u *mongoUser) DBType() data.DBType {
	return mongo.DBType
}
