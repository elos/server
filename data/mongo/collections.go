package mongo

import (
	"fmt"
	"github.com/elos/server/data"
	"gopkg.in/mgo.v2"
)

const DefaultDatabase string = "test"

var Collections = map[data.Kind]string{
	"user":  "users",
	"event": "events",
}

func database(s *mgo.Session) *mgo.Database {
	return s.DB(DefaultDatabase)
}

func collectionFor(s *mgo.Session, m data.Model) (*mgo.Collection, error) {
	collectionForKind := Collections[m.Kind()]

	if collectionForKind == "" {
		err := fmt.Errorf("No collection name has been specified for the model type %s", m.Kind())
		return nil, err
	}

	return database(s).C(collectionForKind), nil
}
