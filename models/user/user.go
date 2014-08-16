package user

import (
	"fmt"
	"time"

	"github.com/elos/server/db"
	"github.com/elos/server/util"
	"gopkg.in/mgo.v2/bson"
)

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

func (u *User) SetId(id bson.ObjectId) {
	u.Id = id
}

func (u *User) GetId() bson.ObjectId {
	return u.Id
}

func (u *User) Save() error {
	err := db.Save(u)

	if err == nil {
		//u.DidSave()
	}

	return err
}

func (u *User) Concerned() []bson.ObjectId {
	a := make([]bson.ObjectId, 1)
	a[0] = u.Id
	return a
}

func (u *User) Kind() db.Kind {
	return Kind
}

func (u *User) EventIdsHash() map[bson.ObjectId]bool {
	hash := make(map[bson.ObjectId]bool, len(u.EventIds))

	for _, id := range u.EventIds {
		hash[id] = true
	}

	return hash
}

func (u *User) Link(property string, m db.Model) {
	switch property {
	case "event":
		if u.EventIdsHash()[m.GetId()] {
			return
			// return nil
		}

		eventId := m.GetId()

		if u.EventIdsHash()[eventId] {
			// return nil
			return
		}

		u.EventIds = append(u.EventIds, eventId)

		m.Link("user", u)

		u.Save()

	default:
		return
	}
}

func FindUserBy(field string, value interface{}) (db.Model, error) {
	user := &User{}

	session := db.NewSession()
	defer session.Close()

	if err := db.CollectionFor(session, user).Find(bson.M{field: value}).One(user); err != nil {
		return user, err
	}

	return user, nil
}

func New() *User {
	return &User{}
}

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
