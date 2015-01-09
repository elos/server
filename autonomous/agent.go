package autonomous

import (
	"github.com/elos/server/data"
	"time"
)

type Agent interface {
	SetDataAgent(data.Agent)
	GetDataAgent() data.Agent
	SetManager(Manager)
	GetManager() Manager

	Start()
	Stop()
	Kill()

	Alive() bool
}

type NewAgent func(db data.DB, a data.Agent, d time.Duration) Agent
