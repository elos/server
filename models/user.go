package models

import (
	"encoding/json"
	"time"

	"github.com/elos/server/db"
	"github.com/elos/server/util"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name      string        `json:"name"`
	Salt      string        `json:"salt"`
	Key       string        `json:"key"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
}

func (u *User) ToJson() ([]byte, error) {
	return json.MarshalIndent(*u, "", "    ")
}

func CreateUser(name string) (User, error) {
	session := db.NewSession()
	defer session.Close()

	user := User{
		Name:      name,
		Salt:      util.RandomString(12),
		Key:       util.RandomString(64),
		CreatedAt: time.Now(),
	}

	usersCollection := session.DB("test").C("users")

	if err := usersCollection.Insert(user); err != nil {
		return user, err
	}

	if err := usersCollection.Find(bson.M{"salt": user.Salt}).One(&user); err != nil {
		return user, err
	}

	return user, nil
}
