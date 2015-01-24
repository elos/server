package transfer

import (
	"github.com/elos/server/conn"
	"github.com/elos/server/data"
	"github.com/elos/server/data/models/serialization"
	// "github.com/elos/server/util"
)

/*
	Takes a well-formed envelope, a database and a connection
	and attempts to remove that record from the database.

	Successful removal prompts a direct data.DELETE response

	Unsuccessful removal prompts a direct data.POST response
	containing the record in question
*/
func DeleteHandler(e *data.Envelope, db data.DB, c conn.Connection) {
	var (
		kind data.Kind
		info map[string]interface{}
	)

	for kind, info = range e.Data {
		model, err := serialization.ModelFor(kind)

		if err != nil { // Unrecognized Type
			c.WriteJSON(data.NewEnvelope(data.POST, map[data.Kind]map[string]interface{}{kind: info}))
			continue
		}

		if err := serialization.PopulateModel(model, &info); err != nil {
			c.WriteJSON(data.NewEnvelope(data.POST, map[data.Kind]map[string]interface{}{kind: info}))
			continue
		}

		if err = db.Delete(model); err != nil {
			c.WriteJSON(data.NewEnvelope(data.POST, map[data.Kind]map[string]interface{}{kind: info}))
			continue
		}

		c.WriteJSON(data.NewPackage(data.DELETE, data.Map(model)))
	}
}
