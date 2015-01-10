package managers

import (
	"github.com/elos/server/autonomous"
	"sync"
)

type NullHub struct {
	Alive            bool
	m                sync.Mutex
	RegisteredAgents map[autonomous.Agent]bool
}

func NewNullHub() *NullHub {
	return &NullHub{
		RegisteredAgents: make(map[autonomous.Agent]bool),
	}
}

func (h *NullHub) StartAgent(a autonomous.Agent) {
	h.m.Lock()
	defer h.m.Unlock()

	h.RegisteredAgents[a] = true
}

func (h *NullHub) StopAgent(a autonomous.Agent) {
	h.m.Lock()
	defer h.m.Unlock()

	h.RegisteredAgents[a] = false
}

func (h *NullHub) Run() {
	h.m.Lock()
	defer h.m.Unlock()

	h.Alive = true
}

func (h *NullHub) Die() {
	h.m.Lock()
	defer h.m.Unlock()

	h.Alive = false
}

func (h *NullHub) Reset() {
	h.m.Lock()
	defer h.m.Unlock()

	h.Alive = false
	h.RegisteredAgents = make(map[autonomous.Agent]bool)
}
