package config

import (
	"github.com/elos/server/data"
	"github.com/elos/server/data/models/user"
	"github.com/elos/server/services"
	// 	"log"
)

var Outfitter *services.Outfitter

func SetupServices(db data.DB) {
	Outfitter = services.NewOutfitter()
	go Outfitter.Run()

	iter, err := db.NewQuery(user.Kind).Execute()
	if err != nil {
	}

	u := user.New()

	for iter.Next(u) {
		services.OutfitUser(Outfitter, db, u)
	}

	if err := iter.Close(); err != nil {
	}

}
