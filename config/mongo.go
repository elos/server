package config

import (
	"log"

	"github.com/elos/server/db"
	"gopkg.in/mgo.v2"
)

func SetupMongo(addr string) *mgo.Session {
	session, err := db.Connect(addr)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Mongo session created")
	}

	return session
}

func ShutdownMongo() {
	db.Close()
}
