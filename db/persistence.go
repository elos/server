package db

import "log"

func Save(m Model) error {
	session := NewSession()
	defer session.Close()

	collection := CollectionFor(session, m)

	// changeInfo, err := ...
	_, err := collection.UpsertId(m.GetId(), m)

	if err != nil {
		log.Printf("Error saving record of kind %s, err: %s", m.Kind(), err)
	}

	return err
}

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

func FindId(m Model) error {
	return PopulateById(m)
}
