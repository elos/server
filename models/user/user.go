package user

import (
	"github.com/elos/data"
	"github.com/elos/data/mongo"
	"github.com/elos/server/models"
	"github.com/elos/server/util"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Returns a new empty user struct
func New( /*db data.DB*/ ) models.User {
	/*
		switch db.Type() {
		case mongo.DBType:
			return &MongoUser{}
		default:
			return &MongoUser{}
		}
	*/
	return &MongoUser{}
}

// Creates a with a NAME
func Create(db data.DB, name string) (models.User, error) {
	user := New()
	user.SetID(mongo.NewObjectID().(bson.ObjectId))
	user.SetCreatedAt(time.Now())
	user.SetName(name)
	user.SetKey(util.RandomString(64))

	if err := db.Save(user); err != nil {
		return user, err
	} else {
		return user, nil
	}
}
