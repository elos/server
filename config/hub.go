package config

import (
	"log"

	"github.com/elos/server/hub"
)

func SetupHub() {
	hub.PrimaryHub = hub.CreateHub()
	go hub.PrimaryHub.Run()
}

func ShutdownHub() {
	log.Printf("ShutdownHub has not yet been implemented")
}
