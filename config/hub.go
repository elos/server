package config

import (
	"github.com/elos/autonomous"
	"github.com/elos/server/managers"
)

var ClientDataHub autonomous.Manager

func SetupClientDataHub() {
	ClientDataHub = managers.NewAgentHub()
	go ClientDataHub.Run()
}

func ShutdownClientDataHub() {
	ClientDataHub.Stop()
}
