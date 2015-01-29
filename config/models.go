package config

import (
	"github.com/elos/data"
	"github.com/elos/server/models/event"
	"github.com/elos/server/models/user"
	"log"
)

const UserKind data.Kind = "user"
const EventKind data.Kind = "event"

var RMap data.RelationshipMap = map[data.Kind]map[data.Kind]data.LinkKind{
	UserKind: {
		EventKind: data.MulLink,
	},
	EventKind: {
		UserKind: data.OneLink,
	},
}

const DataVersion = 1

func (s *Server) SetupModels() data.Schema {
	sch, err := data.NewSchema(&RMap, DataVersion)
	if err != nil {
		log.Fatal(err)
	}

	sch.Register(UserKind, func() data.Model { return user.New() })
	sch.Register(EventKind, func() data.Model { return event.New() })

	user.SetupModel(sch, UserKind, 1)
	event.SetupModel(sch, EventKind, 1)

	return sch
}
