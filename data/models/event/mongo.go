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

	// Properties
	Name      string    `json:"name"`
	StartTime time.Time `json:"start_time" bson:"start_time"`
	EndTime   time.Time `json:"end_time" bson:"end_time"`

	// Links
	// User   Link          `json:"user" bson:"user"`
	UserId data.ID `json:"user_id" bson:"user_id,omitempty"`
}

func (e *MongoEvent) Save() error {
	return db.Save(e)
}

func (e *MongoEvent) Concerned() []data.ID {
	a := make([]data.ID, 1)
	a[0] = e.UserId
	return a
}

func (e *MongoEvent) SetUser(userId data.ID) error {
	if err := data.CheckID(userId); err != nil {
		return err
	}

	if e.UserId == userId {
		return nil
	}

	e.UserId = userId

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
