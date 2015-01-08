package conn

import (
	"github.com/elos/server/data"
)

type NullConnection struct {
	Writes map[interface{}]bool
	Reads  map[interface{}]bool
	Closed bool
	Error  error
	agent  data.Agent
}

func NewNullConnection(a data.Agent) *NullConnection {
	return (&NullConnection{agent: a}).Reset()
}

func (c *NullConnection) Reset() *NullConnection {
	c.Writes = make(map[interface{}]bool)
	c.Reads = make(map[interface{}]bool)
	c.Closed = false
	c.Error = nil
	return c
}

func (c *NullConnection) WriteJSON(v interface{}) error {
	if c.Error != nil {
		return c.Error
	}

	if c.Closed {
		return ConnectionClosedError
	}

	c.Writes[v] = true
	return nil
}

func (c *NullConnection) ReadJSON(v interface{}) error {
	if c.Error != nil {
		return c.Error
	}

	if c.Closed {
		return ConnectionClosedError
	}

	c.Reads[v] = true
	return nil
}

func (c *NullConnection) Close() error {
	if c.Error != nil {
		return c.Error
	}

	c.Closed = true
	return nil
}

func (c *NullConnection) Agent() data.Agent {
	return c.agent
}
