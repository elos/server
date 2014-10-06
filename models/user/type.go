package user

import (
	"fmt"
	"time"

	"github.com/elos/server/db"
	"github.com/elos/server/util"
	"gopkg.in/mgo.v2/bson"
)

// Definition {{{

const Kind db.Kind = "user"

type User struct {
	// Core
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`

	// Properties
	Name string `json:"name"`
	Key  string `json:"key"`

	// Links
	EventIds []bson.ObjectId `json:"event_ids", bson:"event_ids"`
}

// }}}

// Constructors {{{

// Returns a new empty user struct
func New() *User {
	return &User{}
}

// Creates a with a NAME
func Create(name string) (db.Model, error) {
	user := &User{
		Id:        bson.NewObjectId(),
		CreatedAt: time.Now(),
		Name:      name,
		Key:       util.RandomString(64),
	}

	if err := user.Save(); err != nil {
		return nil, err
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
func Authenticate(id string, key string) (db.Model, bool, error) {
	user, err := Find(bson.ObjectIdHex(id))

	if err != nil {
		return user, false, err
	}

	if user.(*User).Key != key {
		return user, false, fmt.Errorf("Invalid key")
	}

	return user, true, nil
}

// Finds a user model by an id
func Find(id bson.ObjectId) (db.Model, error) {
	user := &User{
		Id: id,
	}

	// Find a user that has specified id
	if err := db.FindId(user); err != nil {
		return user, err
	}

	return user, nil
}

// Finds a user by some field and its value
func FindUserBy(field string, value interface{}) (db.Model, error) {
	user := &User{}

	session := db.NewSession()
	defer session.Close()

	if err := db.CollectionFor(session, user).Find(bson.M{field: value}).One(user); err != nil {
		return user, err
	}

	return user, nil
}

// }}}
