package config

import (
	"github.com/elos/server/autonomous/managers"
	"github.com/elos/server/data"
	"github.com/elos/server/data/models/user"
	// 	"log"
)

var Outfitter *managers.Outfitter

func SetupServices(db data.DB) {
	return
	Outfitter = managers.NewOutfitter()
	go Outfitter.Run()

	iter, err := db.NewQuery(user.Kind).Execute()
	if err != nil {
	}

	u := user.New()

	for iter.Next(u) {
		managers.OutfitUser(Outfitter, db, u)
	}

	if err := iter.Close(); err != nil {
	}

}
