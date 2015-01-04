package logging

import (
	"fmt"
	"strings"
)

type Logger interface {
	Log(...interface{})
	Logf(string, ...interface{})
	Logs(string, string)
	Logsf(string, string, ...interface{})
}

type AbstractLogger struct{}

var Log Logger

func SetLog(l Logger) {
	Log = l
}

func FormatService(service string) string {
	upper := strings.ToUpper(service)
	if len(upper) > 6 {
		upper = upper[:6]
	}
	return upper
}

func FormatLogMessage(service string, message string) string {
	return fmt.Sprintf("[%-6s]: %s", FormatService(service), message)
}
