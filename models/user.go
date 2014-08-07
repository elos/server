package models

import (
	"encoding/json"

	"github.com/elos/server/config"
	"github.com/elos/server/util"
	"gopkg.in/mgo.v2/bson"

	"time"
)

type User struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Salt      string
	Key       string
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) ToJson() ([]byte, error) {
	return json.MarshalIndent(*u, "", "    ")
}

func CreateUser() (User, error) {
	session := config.MongoSession.Copy()

	user := User{
		Salt:      util.RandomString(12),
		Key:       util.RandomString(64),
		CreatedAt: time.Now(),
	}

	usersCollection := session.DB("test").C("users")

	if err := usersCollection.Insert(user); err != nil {
		return user, err
	}

	if err := usersCollection.Find(bson.M{"salt": user.Salt}).One(user); err != nil {
		return user, err
	}

	return user, nil
}
