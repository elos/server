package hub

import (
	"github.com/elos/server/models"
	"github.com/elos/server/util"
	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
)

// The primary hub to be used by the server
var PrimaryHub *Hub

func Setup() {
	PrimaryHub = New()
	go PrimaryHub.Run()
}

/*
	A hub maintains a set of connections, and broadcasts
	to those connections
*/
type Hub struct {
	// Registered connections
	Channels map[bson.ObjectId]*HubChannel

	// Channel to register new HubConnections
	Register chan HubConnection

	// Channel to unregister stale/closed HubConnections
	Unregister chan HubConnection
}

func New() *Hub {
	return &Hub{
		Channels:   make(map[bson.ObjectId]*HubChannel),
		Register:   make(chan HubConnection),
		Unregister: make(chan HubConnection),
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
			util.Log("[Hub] Recieved a model from ModelUpdates")
			recipientIds := m.Concerned()

			for _, recipientId := range recipientIds {
				c := h.Channels[recipientId]
				if c != nil {
					c.WriteJson(m)
				}
			}

			util.Log("[Hub] Sent out the updated model")
		}
	}
}

func (h *Hub) FindOrCreateChannel(id bson.ObjectId) *HubChannel {
	// Lookup the channel by id
	_, present := h.Channels[id]

	// If the channel is not present, create it
	if !present {
		h.Channels[id] = &HubChannel{
			Sockets: make([]*websocket.Conn, 0),
			Send:    make(chan []byte),
		}
	}

	// Return the current, or new channel
	return h.Channels[id]
}

func (h *Hub) SendJson(agent Agent, v interface{}) {
	h.Channels[agent.GetId()].WriteJson(v)
}
