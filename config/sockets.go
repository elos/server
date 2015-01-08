package config

import (
	"github.com/elos/server/data"
	"github.com/elos/server/hub"
)

func SetupSockets(db data.DB) {
	hub.Setup(db)
}

func ShutdownSockets() {
	hub.Shutdown()
}
