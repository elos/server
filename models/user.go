package models

import (
	"time"

	"github.com/elos/server/db"
	"github.com/elos/server/util"
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
	err := db.Save(UserKind, u)

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

func (u *User) EventIdsHash() map[bson.ObjectId]bool {
	hash := make(map[bson.ObjectId]bool, len(u.EventIds))

	for _, id := range u.EventIds {
		hash[id] = true
	}

	return hash
}

func (u *User) AddEvent(e *Event) error {
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

func CreateUser(name string) (*User, error) {
	user := User{
		Id:        bson.NewObjectId(),
		CreatedAt: time.Now(),
		Name:      name,
		Key:       util.RandomString(64),
	}

	if err := user.Save(); err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}

func AuthenticateUser(id string, key string) (User, bool, error) {
	user := User{
		Id: bson.ObjectIdHex(id),
	}

	// Find a user that has specified id
	if err := db.FindId(UserKind, &user); err != nil {
		return user, false, err
	}

	// Check if the key matches the supplied one
	if user.Key != key {
		return user, false, nil
	}

	return user, true, nil
}

func FindUserBy(field string, value interface{}) (User, error) {
	user := User{}

	session := db.NewSession()
	defer session.Close()

	if err := db.CollectionFor(session, UserKind).Find(bson.M{field: value}).One(&user); err != nil {
		return user, err
	}

	return user, nil
}
