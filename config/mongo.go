package config

import (
	"log"

	"gopkg.in/mgo.v2"
)

var MongoSession *mgo.Session

func SetupMongo(addr string) *mgo.Session {
	var err error

	if MongoSession, err = mgo.Dial(addr); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Mongo session created")
	}

	return MongoSession
}

func ShutdownMongo() {
	MongoSession.Close()
}
