package sockets

import (
	"github.com/elos/server/util"
)

// Logging {{{

const ServiceName string = "Hub"

func Log(logMessage string) {
	util.Logs(ServiceName, logMessage)
}

func Logf(format string, v ...interface{}) {
	util.Logsf(ServiceName, format, v...)
}

// Logging }}}
