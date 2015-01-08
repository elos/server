package hub

import (
	"github.com/elos/server/services/autonomous"
)

type Hub interface {
	Run()

	StartAgent(autonomous.Agent)
	StopAgent(autonomous.Agent)

	Die()
}
