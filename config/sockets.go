package config

import (
	"github.com/elos/server/data"
	"github.com/elos/server/sockets"
)

func SetupSockets(db data.DB) {
	sockets.Setup(db)
}

func ShutdownSockets() {
	sockets.Shutdown()
}
