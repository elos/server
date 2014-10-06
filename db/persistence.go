package db

import "log"

// Every saved mode is broadcasted over this channel
var ModelUpdates chan Model = make(chan Model)

// Saves a model, broadcasted that save over ModelUpdates
func Save(m Model) error {
	session := NewSession()
	defer session.Close()

	collection := CollectionFor(session, m)

	// changeInfo, err := ...
	_, err := collection.UpsertId(m.GetId(), m)

	if err != nil {
		log.Printf("Error saving record of kind %s, err: %s", m.Kind(), err)
	}

	if err == nil {
		ModelUpdates <- m
	}

	return err
}

// Populates the model data for an empty struct with a specified id
func PopulateById(m Model) error {
	session := NewSession()
	defer session.Close()

	collection := CollectionFor(session, m)

	err := collection.FindId(m.GetId()).One(m)

	if err != nil {
		log.Printf("There was an error populating the %s model, error: %v", m.Kind(), err)
	}

	return err
}

// Alias of PopulateById()
func FindId(m Model) error {
	return PopulateById(m)
}
