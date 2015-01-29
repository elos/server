package config

import (
	"github.com/elos/server/models/event"
	"github.com/elos/server/models/user"
)

func (s *Server) Sandbox() {
	/* free sandbox at the beginning of server,
	nice to test eventual functionality */

	if s.Store == nil {
		return
	}

	u := user.New()
	e := event.New()

	u.SetID(s.NewObjectID())
	e.SetID(s.NewObjectID())

	u.SetName("Sandy Sandbox")
	e.SetName("Sandy's Party")

	e.SetUser(u)

	s.Save(u)
	s.Save(e)

	Logf("User id: %s", u.ID())
	Logf("Event id: %s", e.ID())
}
