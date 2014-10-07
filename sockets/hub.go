package sockets

import (
	"github.com/elos/server/db"
	"github.com/elos/server/models"
	"github.com/elos/server/util"
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
	Register chan *Connection

	// Channel to unregister stale/closed Connections
	Unregister chan *Connection
}

func NewHub() *Hub {
	return &Hub{
		Channels:   make(map[bson.ObjectId]*Channel),
		Register:   make(chan *Connection),
		Unregister: make(chan *Connection),
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
			go h.RegisterConnection(c)
		case c := <-h.Unregister:
			go h.UnregisterConnection(c)
		case m := <-db.ModelUpdates:
			go h.NotifyConcerned(m)
		}
	}
}

/*
	Register a new connection with the hub, adding that connections socket
	to the channel matching that connections Agent ID
*/
func (h *Hub) RegisterConnection(conn *Connection) {
	util.Logf("[Hub] Registering a new socket for agent id %s", conn.Agent.GetId())

	h.FindOrCreateChannel(conn.Agent.GetId()).AddConnection(conn)

	util.Logf("[Hub] New socket registered for agent id %s", conn.Agent.GetId())
}

/*
	Unregister a connection with the hub, removing that connection's socket
	from the channel matching that connection's Agent ID
*/
func (h *Hub) UnregisterConnection(conn *Connection) {
	util.Logf("[Hub] Unregistering a socket for agent id %s", conn.Agent.GetId())

	// Lookup the channel registered for the agent
	channel := h.Channels[conn.Agent.GetId()]

	if channel != nil {
		// Remove the specified socket if the channel exists
		channel.RemoveConnection(conn)
	}

	util.Logf("[Hub] One socket removed for agent id %s", conn.Agent.GetId())
}

func (h *Hub) NotifyConcerned(m db.Model) {
	util.Log("[Hub] Recieved a model from ModelUpdates")

	p := &Package{
		Action: "POST",
		Data:   models.Map(m),
	}

	for _, recipientId := range m.Concerned() {
		h.SendPackage(recipientId, p)
	}

	util.Log("[Hub] Sent out the updated model")
}

func (h *Hub) FindOrCreateChannel(id bson.ObjectId) *Channel {
	// Lookup the channel by id
	_, present := h.Channels[id]

	// If the channel is not present, create it
	if !present {
		h.Channels[id] = &Channel{
			Connections: make([]*Connection, 0),
			Send:        make(chan []byte),
		}
	}

	// Return the current, or new channel
	return h.Channels[id]
}

func (h *Hub) SendPackage(recipientId bson.ObjectId, p *Package) {
	c := h.Channels[recipientId]
	if c != nil {
		c.WriteJSON(p)
	}
}

func (h *Hub) SendJSON(agent Agent, v interface{}) {
	h.Channels[agent.GetId()].WriteJSON(v)
}
