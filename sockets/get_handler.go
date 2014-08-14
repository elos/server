package sockets

import (
	"github.com/elos/server/db"
	"github.com/elos/server/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getHandler(e *Envelope, hc *Connection) {
	for kind, data := range e.Data {
		model, err := Serialize(kind, data)

		if err != nil {
			PrimaryHub.SendJSON(hc.Agent, util.ApiError{400, 400, "Oh shit", ""})
		}

		if id := bson.ObjectIdHex(data["id"].(string)); id != bson.ObjectId("") {
			model.SetId(id)
		}

		err = db.PopulateById(kind, model)

		if err != nil {
			if err == mgo.ErrNotFound {
				// Handle the error here
				PrimaryHub.SendJSON(hc.Agent, util.ApiError{404, 404, "Not Found", "Bad id?"})
			}
			// Otherwise we don't know
			PrimaryHub.SendJSON(hc.Agent, util.ApiError{400, 400, "Oh shit", ""})
		}

		PrimaryHub.SendJSON(hc.Agent, model)
	}
}
