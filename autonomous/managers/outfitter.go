package managers

import (
	"github.com/elos/data"
	"github.com/elos/server/autonomous"
	"github.com/elos/server/autonomous/agents"
	"time"
)

type Outfitter struct {
	StartAgent chan autonomous.Agent
	StopAgent  chan autonomous.Agent

	registeredAgents map[autonomous.Agent]bool
}

func NewOutfitter() *Outfitter {
	return &Outfitter{
		StartAgent:       make(chan autonomous.Agent),
		StopAgent:        make(chan autonomous.Agent),
		registeredAgents: make(map[autonomous.Agent]bool),
	}
}

func (o *Outfitter) GetId() string {
	return "outfitter"
}

func (o *Outfitter) RequestStop(a autonomous.Agent) {
	o.StopAgent <- a
}

func (o *Outfitter) Run() {
	for {
		select {
		case a := <-o.StartAgent:
			go a.Start()
			o.registeredAgents[a] = true
		case a := <-o.StopAgent:
			go a.Stop()
			o.registeredAgents[a] = false
		}
	}
}

var DefaultAgents map[time.Duration]autonomous.NewAgent = map[time.Duration]autonomous.NewAgent{
	agents.DefaultSleepAgentStartPeriod: agents.NewSleepAgent,
}

func OutfitUser(o *Outfitter, db data.DB, a data.Identifiable) {
	for duration, newAgentFunc := range DefaultAgents {
		autonomousAgent := newAgentFunc(db, a, duration)
		o.StartAgent <- autonomousAgent
	}
}
