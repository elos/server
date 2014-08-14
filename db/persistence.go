package db

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"
)

func Save(v Model) error {
	k := v.Kind()
	log.Printf("\nThe database is trying to save a model:\n %#v", v)
	log.Printf("The id the db is tring to save is %s", v.GetId())

	// A database imposed restriction on persistence
	if v.GetId() == bson.ObjectId("") {
		err := fmt.Errorf("Missing Id")
		log.Printf("Error saving record of kind %s, err: %s", k, err)
		return err
	}

	session := NewSession()
	defer session.Close()

	collection := CollectionFor(session, k)

	// changeInfo, err := ...
	_, err := collection.UpsertId(v.GetId(), v)

	if err != nil {
		log.Printf("Error saving record of kind %s, err: %s", k, err)
	}

	return err
}

func PopulateById(v Model) error {
	k := v.Kind()

	session := NewSession()
	defer session.Close()

	collection := CollectionFor(session, k)

	err := collection.FindId(v.GetId()).One(v)

	if err != nil {
		log.Printf("There was an error populating the %s model, error: %v", k, err)
	}

	return err
}

func FindId(v Model) error {
	return PopulateById(v)
}
