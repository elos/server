package routes

import (
	"github.com/elos/server/util"
)

// Logging {{{

const ServiceName string = "Routes"

func Log(logMessage string) {
	util.Logs(ServiceName, logMessage)
}

func Logf(format string, v ...interface{}) {
	util.Logsf(ServiceName, format, v...)
}

// Logging }}}
