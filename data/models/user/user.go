package user

import (
	"github.com/elos/server/data"
	"github.com/elos/server/util"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const Kind data.Kind = "user"

type User interface {
	data.Model
	/*
		SetCreatedAt(time.Time)
		GetCreatedAt(time.Time)
		SetName(string)
		GetName(string)
		SetKey(string)
		GetKey(string)
	*/
}

// Returns a new empty user struct
func New( /*db data.DB*/ ) User {
	/*
		switch db.Type() {
		case "mongo":
			return &MongoUser{}
		default:
			return &MongoUser{}
		}
	*/
	return &MongoUser{}
}

// Creates a with a NAME
func Create(name string) (User, error) {
	user := &MongoUser{
		ID:        data.NewObjectID().(bson.ObjectId),
		CreatedAt: time.Now(),
		Name:      name,
		Key:       util.RandomString(64),
	}

	if err := user.Save(); err != nil {
		return user, err
	} else {
		return user, nil
	}
}
