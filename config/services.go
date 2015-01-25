package config

import (
	"github.com/elos/data"
	"github.com/elos/server/autonomous/managers"
	"github.com/elos/server/models/user"
)

var Outfitter *managers.Outfitter

func SetupServices(db data.DB) {
	Outfitter = managers.NewOutfitter()
	go Outfitter.Run()

	iter, err := db.NewQuery(UserKind).Execute()
	if err != nil {
	}

	u := user.New()

	for iter.Next(u) {
		managers.OutfitUser(Outfitter, db, u)
	}

	if err := iter.Close(); err != nil {
	}

}
