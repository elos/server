package routes

import (
	"github.com/elos/server/util/logging"
)

// The name of this package as a service for the server
const ServiceName string = "Routes"

func log(v ...interface{}) {
	logging.Log.Logs(ServiceName, v...)
}

func logf(format string, v ...interface{}) {
	logging.Log.Logsf(ServiceName, format, v...)
}
