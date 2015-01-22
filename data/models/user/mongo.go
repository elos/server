package user

import (
	"github.com/elos/server/data"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type MongoUser struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`

	// Properties
	Name string `json:"name"`
	Key  string `json:"key"`

	// Links
	EventIds []data.ID `json:"event_ids", bson:"event_ids"`
}

func (u *MongoUser) Kind() data.Kind {
	return Kind
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

func (u *MongoUser) Save() error {
	return db.Save(u)
}

func (u *MongoUser) Concerned() []data.ID {
	a := make([]data.ID, 1)
	a[0] = u.ID
	return a
}
