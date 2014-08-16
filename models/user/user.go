package user

import (
	"fmt"
	"time"

	"github.com/elos/server/db"
	"github.com/elos/server/models"
	"github.com/elos/server/util"
	"gopkg.in/mgo.v2/bson"
)

const Kind db.Kind = models.UserKind

func New() *models.User {
	return &models.User{}
}

func Create(name string) (*models.User, error) {
	user := models.User{
		Id:        bson.NewObjectId(),
		CreatedAt: time.Now(),
		Name:      name,
		Key:       util.RandomString(64),
	}

	if err := user.Save(); err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}

func Authenticate(id string, key string) (*models.User, bool, error) {
	user, err := Find(bson.ObjectIdHex(id))

	if err != nil {
		return user, false, err
	}

	if user.Key != key {
		return user, false, fmt.Errorf("Invalid key")
	}

	return user, true, nil
}

func Find(id bson.ObjectId) (*models.User, error) {
	user := &models.User{
		Id: id,
	}

	// Find a user that has specified id
	if err := db.FindId(user); err != nil {
		return user, err
	}

	return user, nil
}
