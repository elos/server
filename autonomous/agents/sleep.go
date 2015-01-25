package agents

import (
	"log"
	"time"

	"github.com/elos/data"
	"github.com/elos/server/autonomous"
)

var DefaultSleepAgentStartPeriod time.Duration = 10 * time.Second

type SleepAgent struct {
	*BaseAgent

	startPeriod time.Duration
	ticker      *time.Ticker

	DB data.DB
}

func NewSleepAgent(db data.DB, a data.Identifiable, d time.Duration) autonomous.Agent {
	return &SleepAgent{
		BaseAgent:   NewBaseAgent(),
		startPeriod: d,
		DB:          db,
	}
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

func (s *SleepAgent) Go() {
	s.IncrementProcesses()
	log.Print("This is the sleep agent checking in")
	// implement what the sleep agent would actually do
	s.DecrementProcesses()
}
