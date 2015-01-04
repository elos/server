package sockets

import (
	"github.com/elos/server/models/user"
	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
)

/*
	Describes the ability to be registered with the hub.
		- A unique identifier is the only requirement
*/
type Agent interface {
	GetId() bson.ObjectId
}

/*
	A connection is a pair consisting of an agent and that agents
	particular socket connections

	Note: a single agent can have multiple connections because they
	can have multiple sockets open to the server at one time
*/
type Connection struct {
	Agent  Agent
	Socket *websocket.Conn
}

/*
	Creates a new connection and registers the connection
	with the PrimaryHub

	Note that this also begins the server reading from this socket
*/
func NewConnection(agent Agent, socket *websocket.Conn) {
	// Create a new connection wrapper for the agent. socket connection pair
	connection := &Connection{
		Agent:  agent,
		Socket: socket,
	}

	// Register our connection with the hub
	PrimaryHub.Register <- connection

	u, _ := user.Find(agent.GetId())

	PrimaryHub.SendModel(agent, u)

	// Start reading messages from the socket
	go connection.Read()
}

func (conn *Connection) Close() {
	PrimaryHub.Unregister <- conn
	conn.Socket.Close()
}

func (conn *Connection) Read() {
	// When we break our for loop, close the connection
	defer conn.Close()

	// TODO add read limit and deadline
	for {
		var e Envelope

		err := conn.Socket.ReadJSON(&e)

		if err != nil {
			Logf("An error occurred while reading a Connection, err: %s", err)

			/*
				If there was an error break inf. loop.
				Function then completes, and defer is called
			*/
			break
		}

		e.Source = conn

		// Handle the message
		go Route(&e)
	}
}

func (conn *Connection) WriteJSON(v interface{}) error {
	return conn.Socket.WriteJSON(v)
}
