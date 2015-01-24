package config

import (
	"github.com/elos/server/data"
	"github.com/elos/server/data/models"
	"github.com/elos/server/data/models/event"
	"github.com/elos/server/data/models/user"
	"github.com/elos/server/data/schema"
	"log"
)

var RMap schema.RelationshipMap = map[data.Kind]map[data.Kind]schema.LinkKind{
	models.UserKind: {
		models.EventKind: schema.MulLink,
	},
	models.EventKind: {
		models.UserKind: schema.OneLink,
	},
}

const DataVersion = 1

func SetupModels(db data.DB) {
	s, err := schema.NewSchema(&RMap, DataVersion)
	if err != nil {
		log.Fatal(err)
	}

	user.SetupModel(s, 1)
	event.SetupModel(s, 1)
}
