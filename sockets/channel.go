package sockets

/*
	A channel is a group of sockets
*/
type Channel struct {
	Connections []*Connection
	Send        chan []byte
}

func NewChannel() *Channel {
	return &Channel{
		Connections: make([]*Connection, 0),
		Send:        make(chan []byte),
	}
}

// Appends a connection to the list of connections on this channel
func (channel *Channel) AddConnection(conn *Connection) {
	channel.Connections = append(channel.Connections, conn)
}

// Removes a connection from the list of connections managed by this channel
func (channel *Channel) RemoveConnection(conn *Connection) {
	DeleteConnection(channel.Connections, conn)
}

func IndexConnection(slice []*Connection, value *Connection) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func DeleteConnection(slice []*Connection, value *Connection) {
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
