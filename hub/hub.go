package hub

import (
	"log"

	"github.com/elos/server/models"
	"github.com/gorilla/websocket"
)

/*
	A hub maintains a set of connections, and broadcasts
	to those connections
*/
type Hub struct {
	// Registered connections
	Channels map[string]*HubChannel

	Register   chan HubConnection
	Unregister chan HubConnection
}

var PrimaryHub Hub

func CreateHub() Hub {
	return Hub{
		Channels:   make(map[string]*HubChannel),
		Register:   make(chan HubConnection),
		Unregister: make(chan HubConnection),
	}
}

func (h *Hub) FindOrCreateChannel(id string) *HubChannel {
	_, present := h.Channels[id]

	if !present {
		h.Channels[id] = &HubChannel{
			Sockets: make([]*websocket.Conn, 0),
			Send:    make(chan []byte),
		}
	}

	return h.Channels[id]
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.Register:
			log.Print("REGISTER")
			h.FindOrCreateChannel(c.User.Id.String()).AddSocket(c.Socket)
			log.Print(h.Channels)
			log.Print(h.Channels[c.User.Id.String()])
		case c := <-h.Unregister:
			log.Print("UNREGISTER")
			h.Channels[c.User.Id.String()].RemoveSocket(c.Socket)
		}
	}
}

func (h *Hub) SendJson(user models.User, v interface{}) {
	log.Print("SENDJSON IN HUB.GO")
	h.Channels[user.Id.String()].WriteJson(v)
}
