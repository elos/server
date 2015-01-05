package config

import (
	"github.com/elos/server/data"
	"github.com/elos/server/data/models/user"
	"github.com/elos/server/services"
	"gopkg.in/mgo.v2/bson"
	"log"
)

var Outfitter *services.Outfitter

func SetupServices(db data.DB) {
	Outfitter = services.NewOutfitter()
	go Outfitter.Run()
	u, e := user.Find(bson.ObjectIdHex("54aa4b3e02a14bbae5000001"))
	log.Print(Outfitter)
	log.Print(u)
	if e != nil {
		return
	}
	services.OutfitUser(Outfitter, db, u)
}
