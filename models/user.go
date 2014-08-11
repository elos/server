package models

import (
	"time"

	"github.com/elos/server/db"
	"github.com/elos/server/util"
	"gopkg.in/mgo.v2/bson"
)

const UserKind db.Kind = "user"

type User struct {
	// Core
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`

	// Properties
	Name string `json:"name"`
	Key  string `json:"key"`

	// Links
	// FriendId  bson.ObjectId
	// FamilyIds []bson.ObjectId
}

func (u *User) GetId() bson.ObjectId {
	return (*u).Id
}

func (u *User) Save() error {
	return db.Save(UserKind, u)
}

/*
// --- EXPERIMENTAL TRACE BULLETS {{{
func (u *User) Get(k db.Key) db.Property {
	if u.GetId() == bson.ObjectIdHex("") {
		return nil
	}

	if err := db.PopulateById(UserKind, u); err != nil {
		return nil
	}

	switch k {
	case "Id":
		return u.Id
	case "CreatedAt":
		return u.CreatedAt
	case "Name":
		return u.Name
	case "Key":
		return u.Key
	// Example OneLink
	case "Friend":
		if u.FriendId == bson.ObjectIdHex("") {
			return nil
		}

		friend := &User{Id: u.FriendId}

		if err := db.LinkOne(UserKind, friend); err != nil {
			return nil
		}

		return friend
	// Example Many Link
	case "Family":
		users := make([]*User, len(u.FamilyIds))

		for i, id := range u.FamilyIds {
			users[i] = &User{Id: id}
			db.LinkOne(UserKind, users[i])
		}

		return users
	default:
		return nil
	}
}

func (u *User) Set(key db.Key, value db.Property) error {
	switch key {
	case "Name":
		u.Name = value.(string)
	case "Friend":
	default:
		return nil
	}
	return u.Save()
}

// --- }}}
*/

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
