package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

const PrimaryDatabase string = "test"

var Collections = map[Kind]string{
	"user":  "users",
	"event": "events",
}

func database(s *mgo.Session) *mgo.Database {
	return s.DB(PrimaryDatabase)
}

func collectionFor(s *mgo.Session, m Model) (*mgo.Collection, error) {
	collectionForKind := Collections[m.Kind()]

	if collectionForKind == "" {
		err := fmt.Errorf("No collection name has been specified for the model type %s", m.Kind())
		return nil, err
	}

	return database(s).C(collectionForKind), nil
}
