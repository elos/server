package hub

import (
	"log"

	"github.com/elos/server/models"
	"github.com/gorilla/websocket"
)

type HubConnection struct {
	User   models.User
	Socket *websocket.Conn
}

func NewConnection(user models.User, socket *websocket.Conn) {
	log.Printf("NEWCONNECTION")
	connection := HubConnection{
		User:   user,
		Socket: socket,
	}

	log.Printf("hey")

	PrimaryHub.Register <- connection

	log.Printf("there")

	go connection.Read()
}

func (hc *HubConnection) Read() {
	log.Printf("READ")
	defer hc.Close()

	// add read limit and deadline

	for {
		var e Envelope

		err := hc.Socket.ReadJSON(&e)

		if err != nil {
			/*
				If there was an error break inf. loop
				Function then completes, and defer is called
			*/
			break
		}

		go Route(e, *hc)
	}
}

func (hc *HubConnection) Close() {
	PrimaryHub.Unregister <- *hc
	hc.Socket.Close()
}
