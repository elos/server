package conn

import (
	"github.com/elos/server/data"
	"github.com/gorilla/websocket"
)

type GorillaConnection struct {
	conn  *websocket.Conn
	agent data.Agent
}

func NewGorillaConnection(c *websocket.Conn, a data.Agent) Connection {
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

func (c *GorillaConnection) Agent() data.Agent {
	return c.agent
}