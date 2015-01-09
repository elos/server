package user

import (
	"github.com/elos/server/data"
)

func (u *User) Save() error {
	return db.Save(u)
}

func (u *User) Concerned() []data.ID {
	a := make([]data.ID, 1)
	a[0] = u.ID
	return a
}

func (u *User) AddEvent(eventId data.ID) error {
	if err := data.CheckID(eventId); err != nil {
		return err
	}

	if !u.EventIdsHash()[eventId] {
		u.EventIds = append(u.EventIds, eventId)
		return u.Save()
	}

	return nil
}

func (u *User) RemoveEvent(eventId data.ID) error {
	if err := data.CheckID(eventId); err != nil {
		return err
	}

	eventIds := u.EventIdsHash()

	if eventIds[eventId] {
		eventIds[eventId] = false
		ids := make([]data.ID, 0)
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
