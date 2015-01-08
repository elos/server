package hub

import (
	"github.com/elos/server/conn"
	"github.com/elos/server/data"
	"github.com/elos/server/data/models"
	"gopkg.in/mgo.v2/bson"
)

// The primary hub to be used by the server
var PrimaryHub *Hub

func Setup(db data.DB) {
	PrimaryHub = NewHub(db)
	go PrimaryHub.Run()
}

func Shutdown() {
	// close all connections
	PrimaryHub = nil
}

/*
	A hub maintains a set of connections, and broadcasts
	to those connections
*/
type Hub struct {
	// Registered channels
	Channels map[bson.ObjectId]*Channel

	// Channel to register new Connections
	Register chan conn.Connection

	// Channel to unregister stale/closed Connections
	Unregister chan conn.Connection

	// The database which the hub listens to for updates
	DB data.DB
}

func NewHub(db data.DB) *Hub {
	return &Hub{
		Channels:   make(map[bson.ObjectId]*Channel),
		Register:   make(chan conn.Connection),
		Unregister: make(chan conn.Connection),
		DB:         db,
	}
}

/*
	Run loop of a hub
	Blocks on register, unregister, and model update channels
*/
func (h *Hub) Run() {
	for {
		select {
		case c := <-h.Register:
			go h.RegisterConnection(c)
		case c := <-h.Unregister:
			go h.UnregisterConnection(c)
		case m := <-*h.DB.GetUpdatesChannel():
			go h.NotifyConcerned(m)
		}
	}
}

/*
	Register a new connection with the hub, adding that connections socket
	to the channel matching that connections Agent ID
*/
func (h *Hub) RegisterConnection(c conn.Connection) {
	id := c.Agent().GetId()

	Logf("Registering a new socket for agent id %s", id)

	h.FindOrCreateChannel(id).AddConnection(c)

	Logf("New socket registered for agent id %s", id)
}

/*
	Unregister a connection with the hub, removing that connection's socket
	from the channel matching that connection's Agent ID
*/
func (h *Hub) UnregisterConnection(c conn.Connection) {
	id := c.Agent().GetId()

	Logf("Unregistering a socket for agent id %s", id)

	// Lookup the channel registered for the agent
	channel := h.Channels[id]

	if channel != nil {
		// Remove the specified socket if the channel exists
		channel.RemoveConnection(c)
	}

	Logf("One socket removed for agent id %s", id)
}

func (h *Hub) NotifyConcerned(m data.Model) {
	Log("Recieved a model from ModelUpdates")

	p := &Package{
		Action: "POST",
		Data:   models.Map(m),
	}

	for _, recipientId := range m.Concerned() {
		h.SendPackage(recipientId, p)
	}

	Log("Sent out the updated model")
}

func (h *Hub) FindOrCreateChannel(id bson.ObjectId) *Channel {
	// Lookup the channel by id
	_, present := h.Channels[id]

	// If the channel is not present, create it
	if !present {
		h.Channels[id] = NewChannel()
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

func (h *Hub) SendJSON(agent data.Agent, v interface{}) {
	h.Channels[agent.GetId()].WriteJSON(v)
}

func (h *Hub) SendModel(agent data.Agent, model data.Model) {
	p := &Package{
		Action: "POST",
		Data:   models.Map(model),
	}

	h.SendPackage(agent.GetId(), p)
}
