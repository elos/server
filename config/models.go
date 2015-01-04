package config

import (
	"github.com/elos/server/db"
	"github.com/elos/server/models/event"
	"github.com/elos/server/models/user"
)

func SetupModels(db db.DB) {
	user.DB = db
	event.DB = db
}
