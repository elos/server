package serialization

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/elos/server/data"
	"github.com/elos/server/data/models"
	"github.com/elos/server/data/models/event"
	"github.com/elos/server/data/models/user"
)

func NewPackage(action string, m data.Model) *data.Package {
	return &data.Package{
		Action: action,
		Data:   Map(m),
	}
}

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
	case models.UserKind:
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
		log.Printf("This is the error: %s", err)
		return err
	}

	if err := json.Unmarshal(bytes, model); err != nil {
		log.Printf("This is the error: %s", err)
		return err
	}

	return nil
}
