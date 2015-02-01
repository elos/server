package config

import (
	"log"

	"github.com/elos/data"
	"github.com/elos/data/mongo"
	"github.com/elos/server/models/event"
	"github.com/elos/server/models/task"
	"github.com/elos/server/models/user"
)

const DataVersion = 1
const UserKind data.Kind = "user"
const EventKind data.Kind = "event"
const TaskKind data.Kind = "task"

var RMap data.RelationshipMap = data.RelationshipMap{
	UserKind: {
		EventKind: data.Link{
			Name: "events",
			Kind: data.MulLink,
		},
		TaskKind: data.Link{
			Name: "tasks",
			Kind: data.MulLink,
		},
	},
	EventKind: {
		UserKind: data.Link{
			Name: "user",
			Kind: data.OneLink,
		},
	},
	TaskKind: {
		UserKind: data.Link{
			Name: "user",
			Kind: data.OneLink,
		},
		TaskKind: data.Link{
			Name: "dependencies",
			Kind: data.MulLink,
		},
	},
}

func (s *Server) SetupStore(addr string) {
	mongo.RegisterKind(UserKind, "users")
	mongo.RegisterKind(EventKind, "events")
	mongo.RegisterKind(TaskKind, "tasks")

	db, err := mongo.NewDB(addr)
	if err != nil {
		log.Fatal(err)
	}

	Log("Database connection established")

	sch, err := data.NewSchema(&RMap, DataVersion)
	if err != nil {
		log.Fatal(err)
	}

	Log("Schema successfully validated")

	s.Store = data.NewStore(db, sch)

	s.Store.Register(UserKind, user.NewM)
	s.Store.Register(EventKind, event.NewM)
	s.Store.Register(TaskKind, task.NewM)

	user.Setup(sch, UserKind, 1)
	event.Setup(sch, EventKind, 1)
	task.Setup(sch, TaskKind, 1)
}
