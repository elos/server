package models

import (
	"encoding/json"
	"time"

	"github.com/elos/server/db"
	"github.com/elos/server/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name      string        `json:"name"`
	Key       string        `json:"key"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
}

func UsersCollection(s *mgo.Session) *mgo.Collection {
	return s.DB("test").C("users")
}

func (u *User) ToJson() ([]byte, error) {
	return json.MarshalIndent(*u, "", "    ")
}

func CreateUser(name string) (User, error) {
	session := db.NewSession()
	defer session.Close()

	user := User{
		Name:      name,
		Key:       util.RandomString(64),
		CreatedAt: time.Now(),
	}

	usersCollection := UsersCollection(session)

	if err := usersCollection.Insert(user); err != nil {
		return user, err
	}

	if err := usersCollection.Find(bson.M{"key": user.Key}).One(&user); err != nil {
		return user, err
	}

	return user, nil
}

func AuthenticateUser(id string, key string) (User, bool, error) {

	user, err := FindUserById(id)

	if err != nil {
		return user, false, err
	}

	if user.Key == key {
		return user, true, nil
	} else {
		return user, false, nil
	}

}

func FindUserBy(field string, value interface{}) (User, error) {
	user := User{}

	session := db.NewSession()
	defer session.Close()

	if err := UsersCollection(session).Find(bson.M{field: value}).One(&user); err != nil {
		return user, err
	}

	return user, nil
}

func FindUserById(id string) (User, error) {
	return FindUserBy("_id", bson.ObjectIdHex(id))
}
