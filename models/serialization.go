package models

import (
	"encoding/json"
	"fmt"

	"github.com/elos/server/data"
	"github.com/elos/server/models/event"
	"github.com/elos/server/models/user"
)

/*
	Returns a map like:
	{ user: { Name: "Nick Landolfi"} }
	of form:
	{ <db.Kind>: <db.Model>}
*/
func Map(m data.Model) map[data.Kind]data.Model {
	return map[data.Kind]data.Model{
		m.Kind(): m,
	}
}

/*
	Returns a new allocated model of db.Kind KIND
*/
func Type(kind data.Kind) (data.Model, error) {
	var model data.Model

	switch kind {
	case event.Kind:
		model = event.New()
	case user.Kind:
		model = user.New()
	default:
		return nil, fmt.Errorf("Unrecognized type")
	}

	return model, nil
}

// Alias for Type(db.Kind)
func ModelFor(kind data.Kind) (data.Model, error) {
	return Type(kind)
}

func PopulateModel(model data.Model, attributes *map[string]interface{}) error {
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
