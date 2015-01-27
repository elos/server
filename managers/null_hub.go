package managers

import (
	"github.com/elos/autonomous"
	"sync"
)

type NullHub struct {
	*autonomous.BaseAgent
	m                sync.Mutex
	RegisteredAgents map[autonomous.Agent]bool
}

func NewNullHub() *NullHub {
	return &NullHub{
		BaseAgent:        autonomous.NewBaseAgent(),
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

	delete(h.RegisteredAgents, a)
}

func (h *NullHub) Reset() {
	h.m.Lock()
	defer h.m.Unlock()

	/*
		h.Alive = false
		h.RegisteredAgents = make(map[autonomous.Agent]bool)
	*/
}
