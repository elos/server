package user

import (
	"github.com/elos/server/data"
)

func (u *MongoUser) EventIdsHash() map[data.ID]bool {
	hash := make(map[data.ID]bool, len(u.EventIds))

	for _, id := range u.EventIds {
		hash[id] = true
	}

	return hash
}
