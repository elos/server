package routes

import (
	"github.com/elos/server/util/logging"
)

const ServiceName string = "Routes"

func Log(logMessage string) {
	logging.Log.Logs(ServiceName, logMessage)
}

func Logf(format string, v ...interface{}) {
	logging.Log.Logsf(ServiceName, format, v...)
}
