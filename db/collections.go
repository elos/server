package db

import (
	"gopkg.in/mgo.v2"
)

const PrimaryDatabase = "test"

var Collections = map[Kind]string{
	"user": "users",
}

func Database(s *mgo.Session) *mgo.Database {
	return s.DB(PrimaryDatabase)
}

func CollectionFor(s *mgo.Session, kind Kind) *mgo.Collection {
	return Database(s).C(Collections[kind])
}
