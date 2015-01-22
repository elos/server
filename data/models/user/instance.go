package user

import (
	"github.com/elos/server/data"
)

func (u *MongoUser) AddEvent(eventId data.ID) error {
	if err := data.CheckID(eventId); err != nil {
		return err
	}

	if !u.EventIdsHash()[eventId] {
		u.EventIds = append(u.EventIds, eventId)
		return u.Save()
	}

	return nil
}

func (u *MongoUser) RemoveEvent(eventId data.ID) error {
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
