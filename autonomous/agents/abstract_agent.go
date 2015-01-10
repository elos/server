package agents

import (
	"github.com/elos/server/autonomous"
	"github.com/elos/server/data"
	"sync"
)

func NewBaseAgent() *BaseAgent {
	return &BaseAgent{
		stop: make(chan bool),
	}
}

type BaseAgent struct {
	running bool
	stop    chan bool

	dataAgent data.Agent
	manager   autonomous.Manager
	processes int

	m sync.Mutex
}

func (b *BaseAgent) SetDataAgent(a data.Agent) {
	b.m.Lock()
	defer b.m.Unlock()
	b.dataAgent = a
}

func (b *BaseAgent) GetDataAgent() data.Agent {
	b.m.Lock()
	defer b.m.Unlock()

	return b.dataAgent
}

func (b *BaseAgent) SetManager(m autonomous.Manager) {
	b.m.Lock()
	defer b.m.Unlock()
	b.manager = m
}

func (b *BaseAgent) GetManager() autonomous.Manager {
	b.m.Lock()
	defer b.m.Unlock()

	return b.manager
}

func (b *BaseAgent) Stop() {

	go func() { b.stop <- true }()
}

func (b *BaseAgent) Kill() {
	// non-blocking
	go func() { b.stop <- true }()
}

func (b *BaseAgent) Alive() bool {
	b.m.Lock()
	defer b.m.Unlock()
	return b.running
}

func (b *BaseAgent) IncrementProcesses() {
	b.m.Lock()
	defer b.m.Unlock()

	b.processes += 1
}

func (b *BaseAgent) DecrementProcesses() {
	b.m.Lock()
	defer b.m.Unlock()

	b.processes -= 1
}
