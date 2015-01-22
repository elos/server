package user

import (
	"fmt"
	"time"

	"github.com/elos/server/data"
	"github.com/elos/server/util"
	"gopkg.in/mgo.v2/bson"
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

// Constructors {{{

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

// }}}

// Type Methods {{{

/*
	Authenticates a user, returning a populated user model
	If the second return value is true, the user's credentials have been validated
	otherwise, the user's credentials were malformed.
*/
func Authenticate(id string, key string) (data.Model, bool, error) {
	user, err := Find(data.NewObjectIDFromHex(id))

	if err != nil {
		return user, false, err
	}

	if user.(*MongoUser).Key != key {
		return user, false, fmt.Errorf("Invalid key")
	}

	return user, true, nil
}

// Finds a user model by an id
func Find(id data.ID) (data.Model, error) {
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
func FindUserBy(field string, value interface{}) (data.Model, error) {
	user := &MongoUser{}

	if err := db.PopulateByField(field, value, user); err != nil {
		return user, err
	}

	return user, nil
}

// }}}
