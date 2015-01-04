package db

import (
	"github.com/elos/server/util/logging"
)

const ServiceName string = "DB"

func Log(v ...interface{}) {
	logging.Log.Logs(ServiceName, v...)
}

func Logf(format string, v ...interface{}) {
	logging.Log.Logsf(ServiceName, format, v...)
}
