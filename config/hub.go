package config

import "github.com/elos/server/hub"

func SetupHub() {
	hub.PrimaryHub = hub.CreateHub()
	go hub.PrimaryHub.Run()
}
