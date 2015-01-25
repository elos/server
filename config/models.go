package config

import (
	"github.com/elos/data"
	"github.com/elos/schema"
	"github.com/elos/server/models"
	"github.com/elos/server/models/event"
	"github.com/elos/server/models/user"
	"log"
)

const UserKind data.Kind = "user"
const EventKind data.Kind = "event"

var RMap schema.RelationshipMap = map[data.Kind]map[data.Kind]schema.LinkKind{
	UserKind: {
		EventKind: schema.MulLink,
	},
	EventKind: {
		UserKind: schema.OneLink,
	},
}

const DataVersion = 1

func SetupModels(db data.DB) {
	s, err := schema.NewSchema(&RMap, DataVersion)
	if err != nil {
		log.Fatal(err)
	}

	user.SetupModel(s, UserKind, 1)
	event.SetupModel(s, EventKind, 1)

	models.Register(UserKind, func() schema.Model { return user.New() })
	models.Register(EventKind, func() schema.Model { return event.New() })
}
