package sockets

import "github.com/gorilla/websocket"

/*
	A channel is a group of sockets
*/
type Channel struct {
	Sockets []*websocket.Conn
	Send    chan []byte
}

// Appends a socket to the list of sockets managed by this channel
func (c *Channel) AddSocket(s *websocket.Conn) {
	c.Sockets = append(c.Sockets, s)
}

// Removes a socket from the list of sockets managed by this channel
func (c *Channel) RemoveSocket(s *websocket.Conn) {
	DeleteSocket(c.Sockets, s)
}

// Indexes a value of type *websocketConn in a slice of type *websocketConn
func IndexSocket(slice []*websocket.Conn, value *websocket.Conn) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

// Deletes a value of type *websocketConn from a slice of type *websocketConn
func DeleteSocket(slice []*websocket.Conn, value *websocket.Conn) {
	i := IndexSocket(slice, value)
	a := slice
	a = a[:i+copy(a[i:], a[i+1:])]
}

/*
	Writes some interface of json values to each socket
	subscribed to this channel
*/
func (c *Channel) WriteJSON(v interface{}) []error {
	var errs []error

	for _, socket := range c.Sockets {
		err := socket.WriteJSON(v)

		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}
