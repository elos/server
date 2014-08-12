package db

import (
	"log"

	"gopkg.in/mgo.v2"
)

const PrimaryDatabase = "test"

var Collections = map[Kind]string{
	"user":  "users",
	"event": "events",
}

func Database(s *mgo.Session) *mgo.Database {
	return s.DB(PrimaryDatabase)
}

func CollectionFor(s *mgo.Session, kind Kind) *mgo.Collection {
	collectionForKind := Collections[kind]

	if collectionForKind == "" {
		log.Printf("No collection name has been specified for the model type %s", kind)
	}

	return Database(s).C(Collections[kind])
}
