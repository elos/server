package config

import (
	"github.com/elos/server/hub"
	"github.com/elos/server/routes"
)

var Verbose bool

func SetVerbosity(verbose bool) {
	Verbose = verbose
	routes.Verbose = &Verbose
	hub.Verbose = &Verbose
}
