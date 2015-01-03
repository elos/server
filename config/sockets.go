package config

import (
	"github.com/elos/server/sockets"
)

func SetupSockets() {
	sockets.Setup()
}

func ShutdownSockets() {
	sockets.Shutdown()
}
