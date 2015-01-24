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

/*
	Returns a new allocated model of db.Kind KIND
*/
func Type(kind data.Kind) (data.Record, error) {
	var model data.Record

	switch kind {
	case models.EventKind:
		model = event.New()
	case models.UserKind:
		model = user.New()
	default:
		return nil, fmt.Errorf("Unrecognized type")
	}

	return model, nil
}

// Alias for Type(db.Kind)
func ModelFor(kind data.Kind) (data.Record, error) {
	return Type(kind)
}

func PopulateModel(model data.Record, attributes *map[string]interface{}) error {
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
