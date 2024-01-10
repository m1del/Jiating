package loggers

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type CustomLogger struct {
	*log.Logger
	prefix string
}

var (
	Info        *CustomLogger
	Error       *CustomLogger
	Debug       *CustomLogger
	Performance *CustomLogger
)

func NewCustomLogger(out *os.File, prefix string, flag int) *CustomLogger {
	return &CustomLogger{
		Logger: log.New(out, "", flag),
		prefix: prefix,
	}
}

func (l *CustomLogger) Output(calldepth int, s string) error {
	return l.Logger.Output(calldepth+1, fmt.Sprintf("%s %s %s", getESTTime(), l.prefix, s))
}

// Overridden methods
func (l *CustomLogger) Printf(format string, v ...interface{}) {
	l.Output(2, fmt.Sprintf(format, v...))
}

func (l *CustomLogger) Print(v ...interface{}) {
	l.Output(2, fmt.Sprint(v...))
}

func (l *CustomLogger) Println(v ...interface{}) {
	l.Output(2, fmt.Sprintln(v...))
}

func (l *CustomLogger) Fatalf(format string, v ...interface{}) {
	l.Output(4, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *CustomLogger) Fatal(v ...interface{}) {
	l.Output(4, fmt.Sprint(v...))
	os.Exit(1)
}

func (l *CustomLogger) Errorf(format string, v ...interface{}) {
	l.Output(4, fmt.Sprintf(format, v...))
}

func (l *CustomLogger) Error(v ...interface{}) {
	l.Output(4, fmt.Sprint(v...))
}

func init() {
	logFormat := log.Lmsgprefix

	Info = NewCustomLogger(os.Stdout, "| [INFO]: ", logFormat)
	Error = NewCustomLogger(os.Stderr, "| [ERROR]: ", logFormat|log.Lshortfile)
	Debug = NewCustomLogger(os.Stdout, "| [DEBUG]: ", logFormat)
	Performance = NewCustomLogger(os.Stdout, "| [PERFORMANCE]: ", logFormat)
}

func getESTTime() string {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatal(err)
	}
	return time.Now().In(loc).Format("2006-01-02 15:04:05 MST")
}

func getCallerInfo() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return "???"
	}

	shortFile := file
	if lastIndex := strings.LastIndex(file, "/"); lastIndex > -1 {
		shortFile = file[lastIndex+1:]
	}

	return shortFile + ":" + strconv.Itoa(line) + ": "
}
