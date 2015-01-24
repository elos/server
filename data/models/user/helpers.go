package user

import (
	"fmt"

	"github.com/elos/server/data"
	"github.com/elos/server/data/models"
	"github.com/elos/server/data/mongo"
	"gopkg.in/mgo.v2/bson"
)

/*
	Authenticates a user, returning a populated user model
	If the second return value is true, the user's credentials have been validated
	otherwise, the user's credentials were malformed.
*/
func Authenticate(db data.DB, id string, key string) (data.Record, bool, error) {
	user, err := Find(db, mongo.NewObjectIDFromHex(id))

	if err != nil {
		return user, false, err
	}

	if user.(*MongoUser).Key != key {
		return user, false, fmt.Errorf("Invalid key")
	}

	return user, true, nil
}

// Finds a user model by an id
func Find(db data.DB, id data.ID) (models.User, error) {
	user := &MongoUser{
		ID: id.(bson.ObjectId),
	}

	// Find a user that has specified id
	if err := db.PopulateById(user); err != nil {
		return user, err
	}

	return user, nil
}

// Finds a user by some field and its value
func FindUserBy(db data.DB, field string, value interface{}) (models.User, error) {
	user := &MongoUser{}

	if err := db.PopulateByField(field, value, user); err != nil {
		return user, err
	}

	return user, nil
}
