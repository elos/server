package models

import (
	"time"

	"github.com/elos/server/db"
	"github.com/elos/server/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Model
	Name string `json:"name"`
	Key  string `json:"key"`
}

func UsersCollection(s *mgo.Session) *mgo.Collection {
	return s.DB("test").C("users")
}

func CreateUser(name string) (User, error) {
	session := db.NewSession()
	defer session.Close()

	user := User{
		Name:  name,
		Key:   util.RandomString(64),
		Model: Model{CreatedAt: time.Now()},
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
