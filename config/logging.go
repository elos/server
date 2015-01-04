package config

import (
	"github.com/elos/server/util/logging"
)

func SetupLog(verbose bool) {
	if verbose {
		logging.SetLog(logging.StdOutLog)
	} else {
		logging.SetLog(logging.NullLog)
	}
}

const ServiceName string = "Config"

func Log(v ...interface{}) {
	logging.Log.Logs(ServiceName, v...)
}

func Logf(format string, v ...interface{}) {
	logging.Log.Logsf(ServiceName, format, v...)
}
