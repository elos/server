package event

import (
	"github.com/elos/data"
	"github.com/elos/data/mongo"
	"github.com/elos/schema"
	"github.com/elos/server/models"
	"gopkg.in/mgo.v2/bson"
)

type MongoEvent struct {
	*models.Based `bson:,inline`
	*models.Named `bson:,inline`
	*models.Timed `bson:,inline`
	userID        bson.ObjectId `json:"user_id" bson:"user_id,omitempty"`
}

func (e *MongoEvent) Kind() data.Kind {
	return CurrentEventKind
}

func (e *MongoEvent) Schema() schema.Schema {
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
	a[0] = e.userID
	return a
}

func (e *MongoEvent) SetUser(u models.User) error {
	return e.Schema().Link(e, u)
}

func (e *MongoEvent) LinkOne(r schema.Model) {
	switch r.(type) {
	case models.User:
		e.userID = r.ID().(bson.ObjectId)
	default:
		return
	}
}

func (e *MongoEvent) LinkMul(r schema.Model) {
	return
}

func (e *MongoEvent) UnlinkOne(r schema.Model) {
	switch r.(type) {
	case models.User:
		if e.userID == r.ID() {
			e.userID = *new(bson.ObjectId)
		}
	default:
		return
	}
}

func (e *MongoEvent) UnlinkMul(r schema.Model) {
	return
}
