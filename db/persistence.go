package db

import "log"

func Save(k Kind, v Model) error {
	session := NewSession()
	defer session.Close()

	collection := CollectionFor(session, k)

	// changeInfo, err := ...
	_, err := collection.UpsertId(v.GetId(), v)

	if err != nil {
		log.Print("Error saving record of kind %s", v)
	}

	return err
}

func PopulateById(k Kind, v Model) error {
	session := NewSession()
	defer session.Close()

	collection := CollectionFor(session, k)

	err := collection.FindId(v.GetId()).One(v)

	if err != nil {
		log.Printf("There was an error populating the %s model, error: %v", k, err)
	}

	return err
}

func FindId(k Kind, v Model) error {
	return PopulateById(k, v)
}
