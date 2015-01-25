package config

import (
	"github.com/elos/server/data"
	"github.com/elos/server/models/event"
	"github.com/elos/server/models/user"
)

func Sandbox(db data.DB) {
	/* free sandbox at the beginning of server,
	nice to test eventual functionality */

	u := user.New()
	e := event.New()

	u.SetID(db.NewObjectID())
	e.SetID(db.NewObjectID())

	u.SetName("Sandy Sandbox")
	e.SetName("Sandy's Party")

	e.SetUser(u)

	u.Save(db)
	e.Save(db)

	Logf("User id: %s", u.GetID())
	Logf("Event id: %s", e.GetID())
}
