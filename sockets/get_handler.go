package sockets

import (
	"github.com/elos/server/data/models"
	"github.com/elos/server/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getHandler(e *Envelope) {
	// kind is db.Kind
	// data is map[string]interface{}
	for kind, data := range e.Data {
		model, err := models.ModelFor(kind)

		if err != nil {
			PrimaryHub.SendJSON(e.Source.Agent, util.ApiError{400, 400, "Oh shit", ""})
			return
		}

		err = models.PopulateModel(model, &data)

		if id := model.GetId(); id == bson.ObjectId("") {
			PrimaryHub.SendJSON(e.Source.Agent, util.ApiError{400, 400, "Invalid ID", ""})
			return
		}

		// FIXME: INJECT!
		err = PrimaryHub.DB.PopulateById(model)

		if err != nil {
			if err == mgo.ErrNotFound {
				// Handle the error here
				PrimaryHub.SendJSON(e.Source.Agent, util.ApiError{404, 404, "Not Found", "Bad id?"})
				return
			}
			// Otherwise we don't know
			PrimaryHub.SendJSON(e.Source.Agent, util.ApiError{400, 400, "Oh shit", ""})
			return
		}

		PrimaryHub.SendJSON(e.Source.Agent, model)
	}
}
