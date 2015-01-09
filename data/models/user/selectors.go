package user

import (
	"github.com/elos/server/data"
)

func (u *User) Kind() data.Kind {
	return Kind
}

func (u *User) SetID(id data.ID) {
	u.ID = id
}

func (u *User) GetID() data.ID {
	return u.ID
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) EventIdsHash() map[data.ID]bool {
	hash := make(map[data.ID]bool, len(u.EventIds))

	for _, id := range u.EventIds {
		hash[id] = true
	}

	return hash
}
