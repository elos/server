package sockets

import (
	"encoding/json"

	"github.com/elos/server/models"
	"github.com/elos/server/util"
)

func postHandler(e *Envelope, c *Connection) {
	for kind, data := range e.Data {
		model, err := models.Type(kind)

		if err != nil {
			PrimaryHub.SendJSON(c.Agent, util.ApiError{400, 400, "Oh shit", ""})
		}

		// Cleanest way I know of transforming
		// the data to the model's schema
		bytes, err := json.Marshal(data)
		if err != nil {
			PrimaryHub.SendJSON(c.Agent, util.ApiError{400, 400, "Error re-marshalling json data", "I need to check maself"})
			return
		}
		if err := json.Unmarshal(bytes, model); err != nil {
			PrimaryHub.SendJSON(c.Agent, util.ApiError{400, 400, "Error unmarshalling json into struct for model", "Check yoself"})
			return
		}

		// not needed in new masterless approach
		// model.SetId(bson.NewObjectId())

		err = model.Save()
		if err != nil {
			PrimaryHub.SendJSON(c.Agent, util.ApiError{400, 400, "Error saving the model", "Check yoself"})
			return
		}

		// It will go out out as ModelUpdate too, through concerned
		PrimaryHub.SendJSON(c.Agent, model)
		// fixme: this results in a double send
	}
}
