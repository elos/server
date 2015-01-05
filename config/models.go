package config

import (
	"github.com/elos/server/data"
	"github.com/elos/server/data/models/event"
	"github.com/elos/server/data/models/user"
)

func SetupModels(db data.DB) {
	user.DB = db
	event.DB = db
}
