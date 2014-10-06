package user

import (
	"github.com/elos/server/db"
	"gopkg.in/mgo.v2/bson"
)

func (u *User) Save() error {
	return db.Save(u)
}

func (u *User) Concerned() []bson.ObjectId {
	a := make([]bson.ObjectId, 1)
	a[0] = u.Id
	return a
}

func (u *User) Link(property string, m db.Model) {
	switch property {
	case "event":
		eventId := m.GetId()

		// If the user already has the event model linked, then return
		if u.EventIdsHash()[eventId] {
			return
		}

		u.EventIds = append(u.EventIds, eventId)

		m.Link("user", u)

		u.Save()
	default:
		return
	}
}
