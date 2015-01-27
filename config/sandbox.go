package config

import (
	"encoding/json"
	"github.com/elos/data"
	"github.com/elos/server/models/event"
	"gopkg.in/mgo.v2/bson"
)

func Sandbox(db data.DB) {
	/* free sandbox at the beginning of server,
	nice to test eventual functionality */

	e := event.New()

	e.SetID(bson.ObjectIdHex("54c73e8ddd637c577a000002"))
	db.PopulateByID(e)
	e.SetName("hey")
	Logf("This is the event: %v", e)
	bytes, _ := json.Marshal(e)
	Logf(string(bytes[:]))
	Logf(e.Name())

	Logf("Event id: %s", e.ID())
}
