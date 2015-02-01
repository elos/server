package config

import (
	"log"

	"github.com/elos/data"
	"github.com/elos/server/models/event"
	"github.com/elos/server/models/task"
	"github.com/elos/server/models/user"
)

func (s *Server) Sandbox() {
	/* free sandbox at the beginning of server,
	nice to test eventual functionality */

	if s.Store == nil {
		return
	}

	u, _ := user.Create(s.Store, data.AttrMap{"name": "Sandy Sandbox"})

	e, _ := event.New(s.Store)

	e.SetID(s.NewID())

	u.SetName("Sandy Sandbox")
	e.SetName("Sandy's Party")

	err := e.SetUser(u)
	if err != nil {
		log.Fatal(err)
	}

	t, _ := task.New(s.Store)
	t.SetID(s.NewID())
	t.SetName("Sandy's Parent Task")

	t1, _ := task.New(s.Store)
	t2, _ := task.New(s.Store)

	t1.SetName("Sandy's Child Task 1")
	t2.SetName("Sandy's Child Task 2")
	t1.SetID(s.NewID())
	t2.SetID(s.NewID())

	u.AddTask(t)
	u.AddTask(t1)
	u.AddTask(t2)

	t1.AddDependency(t)
	t2.AddDependency(t)

	s.Save(t)
	s.Save(t1)
	s.Save(t2)

	if err = s.Save(u); err != nil {
		log.Fatal(err)
	}
	if err = s.Save(e); err != nil {
		log.Fatal(err)
	}

	Logf("User id: %s", u.ID())
	Logf("Event id: %s", e.ID())
}
