package sockets

import (
	"github.com/elos/server/util"
	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
)

// Any type that satisfies this interface
// can register with the hub
type Agent interface {
	GetId() bson.ObjectId
}

type Connection struct {
	Agent  Agent
	Socket *websocket.Conn
}

func NewConnection(agent Agent, socket *websocket.Conn) {
	// Create a new connection wrapper for the agent. socket connection pair
	connection := Connection{
		Agent:  agent,
		Socket: socket,
	}

	// Register our connection with the hub
	PrimaryHub.Register <- connection

	// Start reading messages from the socket
	go connection.Read()
}

func (hc *Connection) Read() {
	// When we break our loop, close the connection
	defer hc.Close()

	// TODO add read limit and deadline

	for {
		var e Envelope

		err := hc.Socket.ReadJSON(&e)

		if err != nil {
			util.Logf("[Hub] An error occurred while reading a Connection, err: %s", err)

			/*
				If there was an error break inf. loop.
				Function then completes, and defer is called
			*/
			break
		}

		e.Agent = hc.Agent

		// Handle the message
		go Route(&e, hc)
	}
}

func (hc *Connection) Close() {
	PrimaryHub.Unregister <- *hc
	hc.Socket.Close()
}
