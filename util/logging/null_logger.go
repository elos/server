package logging

type NullLogger struct {
	AbstractLogger
}

func (l *NullLogger) Log(v ...interface{}) {}

func (l *NullLogger) Logf(format string, v ...interface{}) {}

func (l *NullLogger) Logs(service string, message string) {}

func (l *NullLogger) Logsf(service string, format string, v ...interface{}) {}

var NullLog *NullLogger = &NullLogger{}
