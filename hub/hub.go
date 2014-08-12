package hub

import (
	"github.com/elos/server/util"
	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
)

/*
	A hub maintains a set of connections, and broadcasts
	to those connections
*/
type Hub struct {
	// Registered connections
	Channels map[*bson.ObjectId]*HubChannel

	// Channel to register new HubConnections
	Register chan HubConnection

	// Channel to unregister stale/closed HubConnections
	Unregister chan HubConnection

	// Channel to update concerned parties
	Update chan *ConcernedUpdate
}

// The primary hub to be used by the server
var PrimaryHub *Hub

func Setup() {
	PrimaryHub = CreateHub()
	go PrimaryHub.Run()
}

func CreateHub() *Hub {
	return &Hub{
		Channels:   make(map[*bson.ObjectId]*HubChannel),
		Register:   make(chan HubConnection),
		Unregister: make(chan HubConnection),
		Update:     make(chan *ConcernedUpdate),
	}
}

func (h *Hub) FindOrCreateChannel(id *bson.ObjectId) *HubChannel {
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

/*
	Run loop of a hub
	Blocks on register and unregister channels
*/
func (h *Hub) Run() {
	for {
		select {
		case c := <-h.Register:
			util.Logf("Hub is registering a new socket for agent id %s", c.Agent.GetId())

			h.FindOrCreateChannel(c.Agent.GetId()).AddSocket(c.Socket)

			util.Logf("New socket registered for agent id %s", c.Agent.GetId())

		case c := <-h.Unregister:
			util.Logf("Hub is UNregistering a new socket for Agent id %s", c.Agent.GetId())

			// Lookup the channel registered for the agent
			channel := h.Channels[c.Agent.GetId()]

			if channel != nil {
				// Remove the specified socket if the channel exists
				channel.RemoveSocket(c.Socket)
			}

			util.Logf("One socket removed for Agent id %s", c.Agent.GetId())
		case p := <-h.Update:
			util.Logf("Hub recieved an update: %v", p)
			c := h.Channels[p.ConcernedId]

			c.WriteJson(p.Package)
			util.Logf("Hub sent an update")
		}
	}
}

func (h *Hub) SendJson(agent Agent, v interface{}) {
	h.Channels[agent.GetId()].WriteJson(v)
}
