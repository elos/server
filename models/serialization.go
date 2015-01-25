package models

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/elos/data"
	"github.com/elos/schema"
)

/*
	Returns a new allocated model of db.Kind KIND
*/
func Type(kind data.Kind) (schema.Model, error) {
	constructor, ok := RegisteredModels[kind]

	if !ok {
		return nil, fmt.Errorf("Unrecognized type")
	}

	return constructor(), nil
}

// Alias for Type(db.Kind)
func ModelFor(kind data.Kind) (schema.Model, error) {
	return Type(kind)
}

func PopulateModel(model schema.Model, attributes *data.AttrMap) error {
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
