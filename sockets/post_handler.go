package sockets

import (
	"encoding/json"
	"log"

	"github.com/elos/server/models"
	"github.com/elos/server/util"
	"gopkg.in/mgo.v2/bson"
)

func postHandler(e *Envelope, hc *Connection) {
	for kind, data := range e.Data {
		model, err := models.Type(kind)

		if err != nil {
			PrimaryHub.SendJSON(hc.Agent, util.ApiError{400, 400, "Oh shit", ""})
		}

		// Cleanest way I know of transforming
		// the data to the model's schema
		bytes, err := json.Marshal(data)
		if err != nil {
			PrimaryHub.SendJSON(hc.Agent, util.ApiError{400, 400, "Error re-marshalling json data", "I need to check maself"})
			return
		}
		if err := json.Unmarshal(bytes, model); err != nil {
			PrimaryHub.SendJSON(hc.Agent, util.ApiError{400, 400, "Error unmarshalling json into struct for model", "Check yoself"})
			return
		}

		model.SetId(bson.NewObjectId())

		log.Printf("Model %#v:", model)
		log.Printf("ID: %s", model.GetId())

		err = model.Save()
		if err != nil {
			PrimaryHub.SendJSON(hc.Agent, util.ApiError{400, 400, "Error saving the model", "Check yoself"})
			return
		}

		// It will go out through the db too, through concerned automaticall
		concerned := model.Concerned()
		concernedMap := make(map[bson.ObjectId]bool, len(concerned))
		for _, c := range concerned {
			concernedMap[c] = true
		}

		// In case this user isn't one of the concerned
		if !concernedMap[hc.Agent.GetId()] {
			PrimaryHub.SendJSON(hc.Agent, model)
		}
	}
}
