package hub

import (
	"github.com/elos/server/util"
	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
)

type Agent interface {
	GetId() *bson.ObjectId
}

type HubConnection struct {
	Agent  Agent
	Socket *websocket.Conn
}

func NewConnection(agent Agent, socket *websocket.Conn) {
	// Create a new connection wrapper for the agent. socket connection pair
	connection := HubConnection{
		Agent:  agent,
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
			util.Logf("An error occurred while reading a HubConnection, err: %s", err)

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
