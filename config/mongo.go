package config

import (
	"log"

	"gopkg.in/mgo.v2"
)

var MongoSession *mgo.Session

func SetupMongo() *mgo.Session {
	var err error

	if MongoSession, err = mgo.Dial("localhost"); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Mongo session %v created", MongoSession)
	}

	return MongoSession
}

func ShutdownMongo() {
	MongoSession.Close()
}
