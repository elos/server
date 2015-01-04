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

func Log(message string) {
	logging.Log.Logs(ServiceName, message)
}

func Logf(format string, v ...interface{}) {
	logging.Log.Logsf(ServiceName, format, v...)
}
