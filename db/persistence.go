package db

import (
	"gopkg.in/mgo.v2/bson"
)

// Saves a model, broadcasted that save over ModelUpdates
func save(db DB, m Model) error {
	session, err := newSession(db)
	if err != nil {
		log(err)
		return err
	}
	defer session.Close()

	collection, err := collectionFor(session, m)
	if err != nil {
		log(err)
		return err
	}

	id := m.GetId()
	if err := CheckId(id); err != nil {
		return err
	}

	// changeInfo, err := ...
	_, err = collection.UpsertId(id, m)

	if err != nil {
		logf("Error saving record of kind %s, err: %s", m.Kind(), err)
	} else {
		*db.GetUpdatesChannel() <- m
	}

	return err
}

// Populates the model data for an empty struct with a specified id
func populateById(db DB, m Model) error {
	session, err := newSession(db)
	if err != nil {
		log(err)
		return err
	}
	defer session.Close()

	collection, err := collectionFor(session, m)
	if err != nil {
		log(err)
		return err
	}

	id := m.GetId()
	if err := CheckId(id); err != nil {
		return err
	}

	err = collection.FindId(m.GetId()).One(m)

	if err != nil {
		logf("There was an error populating the %s model, error: %v", m.Kind(), err)
	}

	return err
}

func populateByField(db DB, m Model, field string, value interface{}) error {
	session, err := newSession(db)
	if err != nil {
		log(err)
		return err
	}
	defer session.Close()

	collection, err := collectionFor(session, m)
	if err != nil {
		log(err)
		return err
	}

	if err := collection.Find(bson.M{field: value}).One(m); err != nil {
		log(err)
		return err
	}

	return nil
}
