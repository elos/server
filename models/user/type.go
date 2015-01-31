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
func New(s data.Store) (models.User, error) {
	switch s.Type() {
	case mongo.DBType:
		return &mongoUser{}, nil
	default:
		return nil, data.ErrInvalidDBType
	}
}

// Creates a with a NAME
func Create(s data.Store, a data.AttrMap) (models.User, error) {
	user, err := New(s)
	if err != nil {
		return user, err
	}

	if id, ok := a["id"].(bson.ObjectId); ok {
		user.SetID(id)
	} else {
		user.SetID(mongo.NewObjectID().(bson.ObjectId))
	}

	if ca, ok := a["created_at"].(time.Time); ok {
		user.SetCreatedAt(ca)
	} else {
		user.SetCreatedAt(time.Now())
	}

	if n, ok := a["name"].(string); ok {
		user.SetName(n)
	}

	user.SetKey(util.RandomString(64))

	if err := s.Save(user); err != nil {
		return user, err
	} else {
		return user, nil
	}
}
