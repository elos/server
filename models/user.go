package models

import (
	"time"

	"github.com/elos/server/db"
	"gopkg.in/mgo.v2/bson"
)

// --- Definition {{{

const UserKind db.Kind = "user"

type User struct {
	// Core
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`

	// Properties
	Name string `json:"name"`
	Key  string `json:"key"`

	// Links
	EventIds []bson.ObjectId `json:"event_ids", bson:"event_ids"`
}

func (u *User) SetId(id bson.ObjectId) {
	u.Id = id
}

func (u *User) GetId() bson.ObjectId {
	return u.Id
}

func (u *User) Save() error {
	err := db.Save(u)

	if err == nil {
		u.DidSave()
	}

	return err
}

// --- }}}

func (u *User) Concerned() []bson.ObjectId {
	a := make([]bson.ObjectId, 1)
	a[0] = u.Id
	return a
}

func (u *User) Kind() db.Kind {
	return UserKind
}

func (u *User) EventIdsHash() map[bson.ObjectId]bool {
	hash := make(map[bson.ObjectId]bool, len(u.EventIds))

	for _, id := range u.EventIds {
		hash[id] = true
	}

	return hash
}

func (u *User) AddEvent(e *Event) error {
	if u.EventIdsHash()[e.Id] {
		return nil
	}

	eventId := e.GetId()

	if u.EventIdsHash()[eventId] {
		return nil
	}

	u.EventIds = append(u.EventIds, eventId)

	if e.UserId != u.Id {
		e.SetUser(u)
	}

	return u.Save()
}

func FindUserBy(field string, value interface{}) (*User, error) {
	user := &User{}

	session := db.NewSession()
	defer session.Close()

	if err := db.CollectionFor(session, user).Find(bson.M{field: value}).One(user); err != nil {
		return user, err
	}

	return user, nil
}

func FindUser(id bson.ObjectId) (*User, error) {
	user := &User{
		Id: id,
	}

	err := db.PopulateById(user)

	return user, err
}
