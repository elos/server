package config

import (
	"github.com/elos/server/autonomous/managers"
	"github.com/elos/server/data"
	"github.com/elos/server/data/models"
	"github.com/elos/server/data/models/user"
)

var Outfitter *managers.Outfitter

func SetupServices(db data.DB) {
	Outfitter = managers.NewOutfitter()
	go Outfitter.Run()

	iter, err := db.NewQuery(models.UserKind).Execute()
	if err != nil {
	}

	u := user.New()

	for iter.Next(u) {
		managers.OutfitUser(Outfitter, db, u)
	}

	if err := iter.Close(); err != nil {
	}

}
