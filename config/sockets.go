package config

import (
	"github.com/elos/server/db"
	"github.com/elos/server/sockets"
)

func SetupSockets(db db.DB) {
	sockets.Setup(db)
}

func ShutdownSockets() {
	sockets.Shutdown()
}
