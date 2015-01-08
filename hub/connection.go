package hub

import (
	"github.com/elos/server/conn"
	"github.com/elos/server/data/models/user"
)

/*
	Creates a new connection and registers the connection
	with the PrimaryHub

	Note that this also begins the server reading from this socket
*/
func NewConnection(c conn.Connection) {
	// Register our connection with the hub
	PrimaryHub.Register <- c

	agent := c.Agent()

	u, _ := user.Find(agent.GetId())

	PrimaryHub.SendModel(agent, u)

	// Start reading messages from the socket
	go ReadConnection(c)
}

func CloseConnection(c conn.Connection) {
	PrimaryHub.Unregister <- c
	c.Close()
}

func ReadConnection(c conn.Connection) {
	// When we break our for loop, close the connection
	defer c.Close()

	// TODO add read limit and deadline
	for {
		var e Envelope

		err := c.ReadJSON(&e)

		if err != nil {
			Logf("An error occurred while reading a Connection, err: %s", err)

			/*
				If there was an error break inf. loop.
				Function then completes, and defer is called
			*/
			break
		}

		e.SourceConnection = c

		// Handle the message
		go Route(&e)
	}
}
