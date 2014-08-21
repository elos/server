package sockets

import (
	"github.com/elos/server/db"
	"github.com/elos/server/models"
	"github.com/elos/server/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getHandler(e *Envelope, c *Connection) {
	for kind, data := range e.Data {
		model, err := models.Type(kind)

		if err != nil {
			PrimaryHub.SendJSON(c.Agent, util.ApiError{400, 400, "Oh shit", ""})
		}

		if id := bson.ObjectIdHex(data["id"].(string)); id != bson.ObjectId("") {
			model.SetId(id)
		}

		err = db.PopulateById(model)

		if err != nil {
			if err == mgo.ErrNotFound {
				// Handle the error here
				PrimaryHub.SendJSON(c.Agent, util.ApiError{404, 404, "Not Found", "Bad id?"})
			}
			// Otherwise we don't know
			PrimaryHub.SendJSON(c.Agent, util.ApiError{400, 400, "Oh shit", ""})
		}

		PrimaryHub.SendJSON(c.Agent, model)
	}
}
