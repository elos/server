package sockets

import (
	"fmt"

	"github.com/elos/server/db"
	"github.com/elos/server/models/event"
	"github.com/elos/server/models/user"
)

func Serialize(kind db.Kind, data map[string]interface{}) (db.Model, error) {
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
