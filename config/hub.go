package config

import (
	"github.com/elos/server/autonomous"
	"github.com/elos/server/autonomous/managers"
)

var ClientDataHub autonomous.Manager

func SetupClientDataHub() {
	ClientDataHub = managers.NewAgentHub()
}

func ShutdownClientDataHub() {
	ClientDataHub.Die()
}
