package conn

import (
	"github.com/elos/data"
)

type GorillaConnection struct {
	conn  AnonConnection
	agent data.Identifiable
}

func NewGorillaConnection(c AnonConnection, a data.Identifiable) Connection {
	return &GorillaConnection{
		conn:  c,
		agent: a,
	}
}

func (c *GorillaConnection) WriteJSON(v interface{}) error {
	return c.conn.WriteJSON(v)
}

func (c *GorillaConnection) ReadJSON(v interface{}) error {
	return c.conn.ReadJSON(v)
}

func (c *GorillaConnection) Close() error {
	return c.conn.Close()
}

func (c *GorillaConnection) Agent() data.Identifiable {
	return c.agent
}
