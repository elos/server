package logging

import (
	"log"
)

type FileLogger struct {
	AbstractLogger
}

func (l *FileLogger) Log(v ...interface{}) {
	log.Print("File Logger Not implemented")
}

func (l *FileLogger) Logf(format string, v ...interface{}) {
	log.Print("File Logger Not implemented")
}

func (l *FileLogger) Logs(service string, v ...interface{}) {
	log.Print("File Logger Not implemented")
}

func (l *FileLogger) Logsf(service string, format string, v ...interface{}) {
	log.Print("File Logger Not implemented")
}

var FileLog *FileLogger = &FileLogger{}
