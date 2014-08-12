package util

import "log"

var Verbose *bool

func Log(v ...interface{}) {
	if *Verbose {
		log.Print(v)
	}
}

func Logf(format string, v ...interface{}) {
	if *Verbose {
		log.Printf(format, v)
	}
}
