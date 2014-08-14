package sockets

import (
	"fmt"

	"github.com/elos/server/db"
	"github.com/elos/server/models"
	"github.com/elos/server/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Envelope struct {
	Agent  Agent
	Action string                  `json:"action"`
	Data   map[db.Kind]interface{} `json:"data"`
}

func Route(e *Envelope, hc *Connection) error {
	switch e.Action {
	case "POST":
		return postHandler(e, hc)
	case "GET":
		return getHandler(e, hc)
	case "DELETE":
		return nil
		// haha, can't delete >:)
	default:
		return fmt.Errorf("Action not recognized")
	}
}

func postHandler(e *Envelope, hc *Connection) error {
	// Echo
	PrimaryHub.SendJSON(hc.Agent, e)
	return nil
}

func getHandler(e *Envelope, hc *Connection) error {
	for kind, data := range e.Data {
		id := bson.ObjectIdHex(data.(map[string]interface{})["id"].(string))
		var (
			err   error
			model db.Model
		)

		util.Logf("[Hub] Id determined: %v", id)

		switch kind {
		case models.UserKind:
			util.Log("[Hub] Determined to be user")
			model = &models.User{
				Id: id,
			}

			err = db.PopulateById(models.UserKind, model)

			util.Logf("[Hub] User looks like %#v", model)
		case models.EventKind:
			util.Log("[Hub] Determined to be event")
			model = &models.Event{
				Id: id,
			}

			err = db.PopulateById(models.EventKind, model)
		default:
			util.Log("[Hub] Determined to be undetermined")
			PrimaryHub.SendJSON(hc.Agent, util.ApiError{400, 400, "Bad", "Stuff"})
		}

		if err != nil {
			if err == mgo.ErrNotFound {
				PrimaryHub.SendJSON(hc.Agent, util.ApiError{404, 404, "Not Found", "Bad id?"})
			}
			PrimaryHub.SendJSON(hc.Agent, util.ApiError{400, 400, "Error!", fmt.Sprintf("Error: %s", e)})
			return err
		}

		PrimaryHub.SendJSON(hc.Agent, model)
	}

	// Echo
	PrimaryHub.SendJSON(hc.Agent, e)
	return nil
}
