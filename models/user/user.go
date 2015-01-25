package user

import (
	"github.com/elos/server/data"
	"github.com/elos/server/models"
	"github.com/elos/server/data/mongo"
	"github.com/elos/server/util"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Returns a new empty user struct
func New( /*db data.DB*/ ) models.User {
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
func Create(db data.DB, name string) (models.User, error) {
	user := &MongoUser{
		ID:        mongo.NewObjectID().(bson.ObjectId),
		CreatedAt: time.Now(),
		Name:      name,
		Key:       util.RandomString(64),
	}

	if err := user.Save(db); err != nil {
		return user, err
	} else {
		return user, nil
	}
}
