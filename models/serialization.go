package models

import (
	"fmt"

	"github.com/elos/server/db"
	"github.com/elos/server/models/event"
	"github.com/elos/server/models/user"
)

func Map(m db.Model) map[db.Kind]db.Model {
	return map[db.Kind]db.Model{
		m.Kind(): m,
	}
}

func Type(kind db.Kind) (db.Model, error) {
	var model db.Model

	switch kind {
	case event.Kind:
		model = user.New()
	case user.Kind:
		model = event.New()
	default:
		return user.New(), fmt.Errorf("Unrecognized type")
	}

	return model, nil
}
