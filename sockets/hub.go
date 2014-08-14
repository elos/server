package sockets

import (
	"github.com/elos/server/db"
	"github.com/elos/server/models"
	"github.com/elos/server/util"
	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
)

// The primary hub to be used by the server
var PrimaryHub *Hub

func Setup() {
	PrimaryHub = NewHub()
	go PrimaryHub.Run()
}

/*
	A hub maintains a set of connections, and broadcasts
	to those connections
*/
type Hub struct {
	// Registered connections
	Channels map[bson.ObjectId]*Channel

	// Channel to register new Connections
	Register chan Connection

	// Channel to unregister stale/closed Connections
	Unregister chan Connection
}

func NewHub() *Hub {
	return &Hub{
		Channels:   make(map[bson.ObjectId]*Channel),
		Register:   make(chan Connection),
		Unregister: make(chan Connection),
	}
}

/*
	Run loop of a hub
	Blocks on register and unregister channels
*/
func (h *Hub) Run() {
	for {
		select {
		case c := <-h.Register:
			util.Logf("[Hub] Registering a new socket for agent id %s", c.Agent.GetId())

			h.FindOrCreateChannel(c.Agent.GetId()).AddSocket(c.Socket)

			util.Logf("[Hub] New socket registered for agent id %s", c.Agent.GetId())
		case c := <-h.Unregister:
			util.Logf("[Hub] Unregistering a socket for agent id %s", c.Agent.GetId())

			// Lookup the channel registered for the agent
			channel := h.Channels[c.Agent.GetId()]

			if channel != nil {
				// Remove the specified socket if the channel exists
				channel.RemoveSocket(c.Socket)
			}

			util.Logf("[Hub] One socket removed for agent id %s", c.Agent.GetId())
		case m := <-models.ModelUpdates:
			p := &Package{
				Action: "POST",
				Data: map[db.Kind]db.Model{
					m.Kind(): m,
				},
			}

			util.Log("[Hub] Recieved a model from ModelUpdates")
			recipientIds := m.Concerned()

			for _, recipientId := range recipientIds {
				c := h.Channels[recipientId]
				if c != nil {
					c.WriteJSON(p)
				}
			}

			util.Log("[Hub] Sent out the updated model")
		}
	}
}

func (h *Hub) FindOrCreateChannel(id bson.ObjectId) *Channel {
	// Lookup the channel by id
	_, present := h.Channels[id]

	// If the channel is not present, create it
	if !present {
		h.Channels[id] = &Channel{
			Sockets: make([]*websocket.Conn, 0),
			Send:    make(chan []byte),
		}
	}

	// Return the current, or new channel
	return h.Channels[id]
}

func (h *Hub) SendJSON(agent Agent, v interface{}) {
	h.Channels[agent.GetId()].WriteJSON(v)
}
