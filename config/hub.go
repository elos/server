package config

import (
	"github.com/elos/server/services/hub"
)

var ClientDataHub hub.Hub

func SetupClientDataHub() {
	ClientDataHub = hub.NewAgentHub()
}

func ShutdownClientDataHub() {
	ClientDataHub.Die()
}
