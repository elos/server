package agents

import (
	"log"
	"time"

	"github.com/elos/server/data"
	"github.com/elos/server/services/autonomous"
)

var DefaultSleepAgentStartPeriod time.Duration = 10 * time.Second

var RunningSleepGos int = 0

type SleepAgent struct {
	running     bool
	stop        chan bool
	startPeriod time.Duration
	ticker      *time.Ticker

	DB        data.DB
	DataOwner data.Agent
	Manager   autonomous.Manager
}

func NewSleepAgent(db data.DB, a data.Agent, d time.Duration) autonomous.Agent {
	return &SleepAgent{
		running:     false,
		stop:        make(chan bool),
		startPeriod: d,
		DB:          db,
		DataOwner:   a,
	}
}

func (s *SleepAgent) SetDataOwner(a data.Agent) {
	if a != nil {
		s.DataOwner = a
	}
}

func (s *SleepAgent) SetManager(m autonomous.Manager) {
	if m != nil {
		s.Manager = m
	}
}

func (s *SleepAgent) GetManager() autonomous.Manager {
	return s.Manager
}

func (s *SleepAgent) Start() {
	s.running = true
	s.ticker = time.NewTicker(s.startPeriod)
	go s.Go()
	for {
		select {
		case _ = <-s.ticker.C:
			go s.Go()
		case _ = <-s.stop:
			s.ticker.Stop()
			s.ticker = nil
			s.running = false
			break
		}
	}
}

func (s *SleepAgent) Stop() {
	s.stop <- true
}

func (s *SleepAgent) Kill() {
	s.stop <- true
}

func (s *SleepAgent) Alive() bool {
	return s.running
}

func (s *SleepAgent) Go() {
	RunningSleepGos = RunningSleepGos + 1
	log.Print("This is the sleep agent checking in")
	// implement what the sleep agent would actually do
	RunningSleepGos = RunningSleepGos - 1
}
