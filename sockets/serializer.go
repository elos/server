package sockets

import (
	"fmt"

	"github.com/elos/server/db"
	"github.com/elos/server/models"
)

func Serialize(kind db.Kind, data map[string]interface{}) (db.Model, error) {
	var model db.Model

	switch kind {
	case models.UserKind:
		model = &models.User{}
	case models.EventKind:
		model = &models.Event{}
	default:
		return &models.User{}, fmt.Errorf("Unrecognized type")
	}

	return model, nil
}
