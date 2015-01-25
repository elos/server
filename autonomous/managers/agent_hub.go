package managers

import (
	"github.com/elos/server/autonomous"
	"log"
)

/*
	A hub maintains a set of connections, and broadcasts
	to those connections
*/
type AgentHub struct {
	Start chan autonomous.Agent
	Stop  chan autonomous.Agent
	die   chan bool

	registeredAgents map[autonomous.Agent]bool
}

func NewAgentHub() *AgentHub {
	return &AgentHub{
		Start:            make(chan autonomous.Agent),
		Stop:             make(chan autonomous.Agent),
		registeredAgents: make(map[autonomous.Agent]bool),
		die:              make(chan bool),
	}
}

func (h *AgentHub) StartAgent(a autonomous.Agent) {
	h.Start <- a
}

func (h *AgentHub) StopAgent(a autonomous.Agent) {
	h.Stop <- a
}

func (h *AgentHub) Run() {
	for {
		select {
		case a := <-h.Start:
			go a.Start()
			h.registeredAgents[a] = true
		case a := <-h.Stop:
			go a.Start()
			delete(h.registeredAgents, a)
		case _ = <-h.die:
			log.Print("hello")
			h.Shutdown()
			log.Print("there")
			break
		}
	}
}

func (h *AgentHub) Die() {
	log.Print("woah")
	h.die <- true
	log.Print("woah")
}

func (h *AgentHub) Shutdown() {
	log.Print("hey")
	for a, _ := range h.registeredAgents {
		log.Print("ho")
		go a.Stop()
		log.Print("ho2")
	}
}
