package db

import (
	"gopkg.in/mgo.v2/bson"
)

// Every saved mode is broadcasted over this channel
var ModelUpdates chan Model = make(chan Model)

// Saves a model, broadcasted that save over ModelUpdates
func Save(m Model) error {
	session, err := newSession()
	if err != nil {
		Log(err)
		return err
	}
	defer session.Close()

	collection, err := CollectionFor(session, m)
	if err != nil {
		Log(err)
		return err
	}

	// changeInfo, err := ...
	_, err = collection.UpsertId(m.GetId(), m)

	if err != nil {
		Logf("Error saving record of kind %s, err: %s", m.Kind(), err)
	} else {
		ModelUpdates <- m
	}

	return err
}

// Populates the model data for an empty struct with a specified id
func PopulateById(m Model) error {
	session, err := newSession()
	if err != nil {
		Log(err)
		return err
	}
	defer session.Close()

	collection, err := CollectionFor(session, m)
	if err != nil {
		Log(err)
		return err
	}

	err = collection.FindId(m.GetId()).One(m)

	if err != nil {
		Logf("There was an error populating the %s model, error: %v", m.Kind(), err)
	}

	return err
}

// Alias of PopulateById()
func FindId(m Model) error { return PopulateById(m) }

func PopulateByField(m Model, field string, value interface{}) error {
	session, err := newSession()
	if err != nil {
		Log(err)
		return err
	}
	defer session.Close()

	collection, err := CollectionFor(session, m)
	if err != nil {
		Log(err)
		return err
	}

	if err := collection.Find(bson.M{field: value}).One(m); err != nil {
		Log(err)
		return err
	}

	return nil
}
