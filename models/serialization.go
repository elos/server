package models

import (
	"encoding/json"
	"fmt"

	"github.com/elos/server/db"
	"github.com/elos/server/models/event"
	"github.com/elos/server/models/user"
)

/*
	Returns a map like:
	{ user: { name: "Nick Landolfi"} }
	of form:
	{ <db.Kind>: <db.Model>}
*/
func Map(m db.Model) map[db.Kind]db.Model {
	return map[db.Kind]db.Model{
		m.Kind(): m,
	}
}

/*
	Returns a new allocated model of db.Kind KIND
*/
func Type(kind db.Kind) (db.Model, error) {
	var model db.Model

	switch kind {
	case event.Kind:
		model = event.New()
	case user.Kind:
		model = user.New()
	default:
		return user.New(), fmt.Errorf("Unrecognized type")
	}

	return model, nil
}

func ModelFor(kind db.Kind) (db.Model, error) {
	return Type(kind)
}

func PopulateModel(model db.Model, attributes *map[string]interface{}) error {
	// Cleanest way I know of transforming the data to the model's schema
	bytes, err := json.Marshal(attributes)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, model); err != nil {
		return err
	}

	return nil
}
