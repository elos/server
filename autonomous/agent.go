package autonomous

import (
	"github.com/elos/server/data"
	"time"
)

type Agent interface {
	SetDataOwner(data.Identifiable)
	GetDataOwner() data.Identifiable
	SetManager(Manager)
	GetManager() Manager

	Start()
	Stop()
	Kill()

	Alive() bool
}

type NewAgent func(db data.DB, a data.Identifiable, d time.Duration) Agent
