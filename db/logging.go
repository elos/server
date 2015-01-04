package db

import (
	"github.com/elos/server/util/logging"
)

const ServiceName string = "DB"

func log(v ...interface{}) {
	logging.Log.Logs(ServiceName, v...)
}

func logf(format string, v ...interface{}) {
	logging.Log.Logsf(ServiceName, format, v...)
}
