package config

import "github.com/elos/server/util"

var Verbose bool

func SetVerbosity(verbose bool) {
	Verbose = verbose
	util.Verbose = &Verbose
}
