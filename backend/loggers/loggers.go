package loggers

import (
	"log"
	"os"
)

var (
	Info        *log.Logger
	Error       *log.Logger
	Debug       *log.Logger
	Performance *log.Logger
)

func init() {
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime)

	// performance
	Performance = log.New(os.Stdout, "PERFORMANCE: ", log.Ldate|log.Ltime)
}
