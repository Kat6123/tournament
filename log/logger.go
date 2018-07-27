package log

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	traceLogger *log.Logger
	debugLogger *log.Logger
	infoLogger  *log.Logger
	errorLogger *log.Logger

	logFlag = log.Ldate | log.Ltime | log.Lshortfile
)

func init() {
	traceLogger = log.New(ioutil.Discard, "Trace: ", logFlag)
	debugLogger = log.New(ioutil.Discard, "Debug: ", logFlag)
	infoLogger = log.New(ioutil.Discard, "Info: ", logFlag)
	errorLogger = log.New(os.Stderr, "Error: ", logFlag)
}

func SetLevel(l Level) {
	switch l {
	case InfoLevel:
		infoLogger.SetOutput(os.Stdout)
	case DebugLevel:
		infoLogger.SetOutput(os.Stdout)
		debugLogger.SetOutput(os.Stdout)
	case TraceLevel:
		infoLogger.SetOutput(os.Stdout)
		debugLogger.SetOutput(os.Stdout)
		traceLogger.SetOutput(os.Stdout)
	}
}

func Trace(format string, v ...interface{}) {
	traceLogger.Printf(format, v...)
}

func Debug(format string, v ...interface{}) {
	debugLogger.Printf(format, v...)
}

func Info(format string, v ...interface{}) {
	infoLogger.Printf(format, v...)
}

func Error(format string, v ...interface{}) {
	errorLogger.Printf(format, v...)
}
