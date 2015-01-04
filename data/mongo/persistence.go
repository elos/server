package mongo

import (
	"github.com/elos/server/data"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Saves a model, broadcasted that save over ModelUpdates
func save(s *mgo.Session, m data.Model) error {
	collection, err := collectionFor(s, m)
	if err != nil {
		// log(err)
		return err
	}

	id := m.GetId()
	if err := data.CheckId(id); err != nil {
		return err
	}

	// changeInfo, err := ...
	_, err = collection.UpsertId(id, m)

	return err
}

// Populates the model data for an empty struct with a specified id
func populateById(s *mgo.Session, m data.Model) error {
	collection, err := collectionFor(s, m)
	if err != nil {
		// log(err)
		return err
	}

	id := m.GetId()
	if err := data.CheckId(id); err != nil {
		return err
	}

	return collection.FindId(m.GetId()).One(m)
}

func populateByField(s *mgo.Session, m data.Model, field string, value interface{}) error {
	collection, err := collectionFor(s, m)
	if err != nil {
		// log(err)
		return err
	}

	if err := collection.Find(bson.M{field: value}).One(m); err != nil {
		// log(err)
		return err
	}

	return nil
}
