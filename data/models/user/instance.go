package user

import (
	"github.com/elos/server/data"
	"gopkg.in/mgo.v2/bson"
)

func (u *User) Save() error {
	return DB.Save(u)
}

func (u *User) Concerned() []bson.ObjectId {
	a := make([]bson.ObjectId, 1)
	a[0] = u.Id
	return a
}

func (u *User) AddEvent(eventId bson.ObjectId) error {
	if err := data.CheckId(eventId); err != nil {
		return err
	}

	if !u.EventIdsHash()[eventId] {
		u.EventIds = append(u.EventIds, eventId)
		return u.Save()
	}

	return nil
}

func (u *User) RemoveEvent(eventId bson.ObjectId) error {
	if err := data.CheckId(eventId); err != nil {
		return err
	}

	eventIds := u.EventIdsHash()

	if eventIds[eventId] {
		eventIds[eventId] = false
		ids := make([]bson.ObjectId, 0)
		for id := range eventIds {
			if eventIds[id] {
				ids = append(ids, id)
			}
		}

		u.EventIds = ids
		return u.Save()
	}

	return nil
}
