package user

import (
	"github.com/elos/server/db"
	"gopkg.in/mgo.v2/bson"
)

func (u *User) Kind() db.Kind {
	return Kind
}

func (u *User) SetId(id bson.ObjectId) {
	u.Id = id
}

func (u *User) GetId() bson.ObjectId {
	return u.Id
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) EventIdsHash() map[bson.ObjectId]bool {
	hash := make(map[bson.ObjectId]bool, len(u.EventIds))

	for _, id := range u.EventIds {
		hash[id] = true
	}

	return hash
}
