package config

import (
	"log"

	"github.com/elos/server/hub"
)

func SetupHub() {
	hub.Setup()
}

func ShutdownHub() {
	log.Printf("ShutdownHub has not yet been implemented")
}
