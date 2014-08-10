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
	// Create a new connection wrapper for the user. socket connection pair
	connection := HubConnection{
		User:   user,
		Socket: socket,
	}

	// Register our connection with the hub
	PrimaryHub.Register <- connection

	// Start reading messages from the socket
	go connection.Read()
}

func (hc *HubConnection) Read() {
	// When we break our loop, close the connection
	defer hc.Close()

	// TODO add read limit and deadline

	for {
		var e Envelope

		err := hc.Socket.ReadJSON(&e)

		if err != nil {
			if Verbose {
				log.Print("An error occurred while reading a HubConnection, err: %s", err)
			}

			/*
				If there was an error break inf. loop
				Function then completes, and defer is called
			*/
			break
		}

		// Handle the message
		go Route(e, *hc)
	}
}

func (hc *HubConnection) Close() {
	PrimaryHub.Unregister <- *hc
	hc.Socket.Close()
}
