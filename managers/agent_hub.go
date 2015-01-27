package managers

import (
	"github.com/elos/autonomous"
	"log"
)

/*
	A hub maintains a set of connections, and broadcasts
	to those connections
*/
type AgentHub struct {
	*autonomous.BaseAgent

	start            chan autonomous.Agent
	stop             chan autonomous.Agent
	registeredAgents map[autonomous.Agent]bool
}

func NewAgentHub() *AgentHub {
	return &AgentHub{
		BaseAgent:        autonomous.NewBaseAgent(),
		start:            make(chan autonomous.Agent),
		stop:             make(chan autonomous.Agent),
		registeredAgents: make(map[autonomous.Agent]bool),
	}
}

func (h *AgentHub) StartAgent(a autonomous.Agent) {
	h.start <- a
}

func (h *AgentHub) StopAgent(a autonomous.Agent) {
	h.stop <- a
}

func (h *AgentHub) Run() {
	h.startup()
	stop := h.BaseAgent.StopChannel()

	for {
		select {
		case a := <-h.start:
			go a.Run()
			h.registeredAgents[a] = true
		case a := <-h.stop:
			go a.Run()
			delete(h.registeredAgents, a)
		case _ = <-*stop:
			log.Print("hello")
			h.shutdown()
			log.Print("there")
			break
		}
	}
}

func (h *AgentHub) startup() {
	h.BaseAgent.Startup()
}

func (h *AgentHub) shutdown() {
	log.Print("hey")
	for a, _ := range h.registeredAgents {
		log.Print("ho")
		go a.Stop()
		log.Print("ho2")
	}
	h.BaseAgent.Shutdown()
}
