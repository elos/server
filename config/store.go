package config

import (
	"log"

	"github.com/elos/data"
	"github.com/elos/data/mongo"
	"github.com/elos/server/models/event"
	"github.com/elos/server/models/user"
)

const DataVersion = 1
const UserKind data.Kind = "user"
const EventKind data.Kind = "event"

var RMap data.RelationshipMap = data.RelationshipMap{
	UserKind: {
		EventKind: data.MulLink,
	},
	EventKind: {
		UserKind: data.OneLink,
	},
}

func (s *Server) SetupStore(addr string) {
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
	s.Store.Register(EventKind, user.NewM)

	user.Setup(sch, UserKind, 1)
	event.Setup(sch, EventKind, 1)
}
