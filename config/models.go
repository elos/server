package config

import (
	"github.com/elos/server/data"
	"github.com/elos/server/models/event"
	"github.com/elos/server/models/user"
)

func SetupModels(db data.DB) {
	user.DB = db
	event.DB = db
}
