package hub

import (
	"github.com/elos/server/conn"
)

/*
	A channel is a group of sockets
*/
type Channel struct {
	Connections []conn.Connection
	Send        chan []byte
}

func NewChannel() *Channel {
	return &Channel{
		Connections: make([]conn.Connection, 0),
		Send:        make(chan []byte),
	}
}

// Appends a connection to the list of connections on this channel
func (channel *Channel) AddConnection(c conn.Connection) {
	channel.Connections = append(channel.Connections, c)
}

// Removes a connection from the list of connections managed by this channel
func (channel *Channel) RemoveConnection(c conn.Connection) {
	DeleteConnection(channel.Connections, c)
}

func IndexConnection(slice []conn.Connection, value conn.Connection) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func DeleteConnection(slice []conn.Connection, value conn.Connection) {
	i := IndexConnection(slice, value)
	a := slice
	a = a[:i+copy(a[i:], a[i+1:])]
}

/*
	Writes some interface of json values to each socket
	subscribed to this channel
*/
func (c *Channel) WriteJSON(v interface{}) []error {
	var errs []error

	for _, conn := range c.Connections {
		err := conn.WriteJSON(v)

		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}
