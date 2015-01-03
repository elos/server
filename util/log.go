package util

import (
	"fmt"
	"log"
)

var Verbose *bool

func Log(v ...interface{}) {
	if *Verbose {
		log.Print(v...)
	}
}

func Logf(format string, v ...interface{}) {
	if *Verbose {
		log.Printf(format, v...)
	}
}

func Logsf(serviceName, format string, v ...interface{}) {
	if *Verbose {
		log.Printf("[%s]: %s", serviceName, fmt.Sprintf(format, v...))
	}
}

func Logs(serviceName string, logMessage string) {
	Logsf(serviceName, logMessage)
}
