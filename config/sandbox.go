package config

import (
	"log"

	"github.com/elos/data"
	"github.com/elos/models/event"
	"github.com/elos/models/task"
	"github.com/elos/models/user"
)

func Sandbox(s data.Store) {
	/* free sandbox at the beginning of server,
	nice to test eventual functionality */

	if s == nil {
		return
	}

	u, _ := user.Create(s, data.AttrMap{"name": "Sandy Sandbox"})

	e, _ := event.New(s)

	e.SetID(s.NewID())

	u.SetName("Sandy Sandbox")
	e.SetName("Sandy's Party")

	err := e.SetUser(u)
	if err != nil {
		log.Fatal(err)
	}

	t, _ := task.New(s)
	t.SetID(s.NewID())
	t.SetName("Sandy's Parent Task")

	t1, _ := task.New(s)
	t2, _ := task.New(s)

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

	log.Printf("User id: %s", u.ID())
	log.Printf("Event id: %s", e.ID())
}
