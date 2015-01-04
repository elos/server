package db

import (
	"gopkg.in/mgo.v2"
)

const PrimaryDatabase string = "test"

var Collections = map[Kind]string{
	"user":  "users",
	"event": "events",
}

func Database(s *mgo.Session) *mgo.Database {
	return s.DB(PrimaryDatabase)
}

func CollectionFor(s *mgo.Session, m Model) *mgo.Collection {
	collectionForKind := Collections[m.Kind()]

	if collectionForKind == "" {
		Logf("No collection name has been specified for the model type %s", m.Kind())
	}

	return Database(s).C(collectionForKind)
}
