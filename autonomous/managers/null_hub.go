package managers

import (
	"github.com/elos/server/autonomous"
	"sync"
)

type NullHub struct {
	m                sync.Mutex
	registeredAgents map[autonomous.Agent]bool
}

func NewNullHub() autonomous.Manager {
	return &NullHub{
		registeredAgents: make(map[autonomous.Agent]bool),
	}
}

func (h *NullHub) StartAgent(a autonomous.Agent) {
	h.m.Lock()
	h.registeredAgents[a] = true
	h.m.Unlock()
}

func (h *NullHub) StopAgent(a autonomous.Agent) {
	h.m.Lock()
	h.registeredAgents[a] = false
	h.m.Unlock()
}

func (h *NullHub) Run() {
}

func (h *NullHub) Die() {
}
