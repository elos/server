package hub

import (
	"log"

	"github.com/elos/server/models"
	"github.com/gorilla/websocket"
)

var Verbose *bool

/*
	A hub maintains a set of connections, and broadcasts
	to those connections
*/
type Hub struct {
	// Registered connections
	Channels map[string]*HubChannel

	// Channel to register new HubConnections
	Register chan HubConnection

	// Channel to unregister stale/closed HubConnections
	Unregister chan HubConnection
}

// The primary hub to be used by the server
var PrimaryHub *Hub

func CreateHub() *Hub {
	return &Hub{
		Channels:   make(map[string]*HubChannel),
		Register:   make(chan HubConnection),
		Unregister: make(chan HubConnection),
	}
}

func (h *Hub) FindOrCreateChannel(id string) *HubChannel {
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
			if *Verbose {
				log.Print("Hub is registering a new socket for User id %s", c.User.Id.String())
			}

			h.FindOrCreateChannel(c.User.Id.String()).AddSocket(c.Socket)

			if *Verbose {
				log.Printf("New socket registered for User id %s", c.User.Id.String())
			}

		case c := <-h.Unregister:
			if *Verbose {
				log.Print("Hub is UNregistering a new socket for User id %s", c.User.Id.String())
			}

			// Lookup the channel registered for the user
			channel := h.Channels[c.User.Id.String()]

			if channel != nil {
				// Remove the specified socket if the channel exists
				channel.RemoveSocket(c.Socket)
			}

			if *Verbose {
				log.Print("One socket removed for User id %s", c.User.Id.String())
			}
		}
	}
}

func (h *Hub) SendJson(user models.User, v interface{}) {
	h.Channels[user.Id.String()].WriteJson(v)
}
