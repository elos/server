package config

import (
	"github.com/elos/server/data"
	"github.com/elos/server/data/models"
	"github.com/elos/server/data/models/event"
	"github.com/elos/server/data/models/schema"
	"github.com/elos/server/data/models/user"
	"log"
)

var SchemaMap schema.SchemaMap = map[data.Kind]map[data.Kind]schema.LinkKind{
	models.UserKind: {
		models.EventKind: schema.MulLink,
	},
	models.EventKind: {
		models.UserKind: schema.OneLink,
	},
}

const DataVersion = 1

func SetupModels(db data.DB) {
	s, err := schema.NewSchema(&SchemaMap, DataVersion)
	if err != nil {
		log.Fatal(err)
	}

	user.SetupModel(s, 1)
	event.SetupModel(s, 1)

	user.SetDB(db)
	event.SetDB(db)
}
