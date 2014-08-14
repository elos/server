package config

import (
	"log"

	"github.com/elos/server/sockets"
)

func SetupSockets() {
	sockets.Setup()
}

func ShutdownSockets() {
	log.Printf("ShutdownHub has not yet been implemented")
}
