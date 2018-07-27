package log

import (
	"flag"
	"fmt"
	"strconv"
)

const (
	TraceLevel Level = iota
	DebugLevel
	InfoLevel
	ErrorLevel
)

// Level is a type of log Level.
type Level int

// String returns string representation of log level.
func (l Level) String() string {
	switch l {
	case TraceLevel:
		return "trace"
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case ErrorLevel:
		return "error"
	default:
		return "unknown log level"
	}
}

// Set parses input string of log level and set value.
func (l *Level) Set(s string) error {
	i, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("parse level %q as int has failed: %v", s, err)
	}

	switch Level(i) {
	case TraceLevel:
		*l = TraceLevel
		return nil
	case DebugLevel:
		*l = DebugLevel
		return nil
	case InfoLevel:
		*l = InfoLevel
		return nil
	case ErrorLevel:
		*l = ErrorLevel
		return nil
	default:
		return fmt.Errorf("wrong level type %q", i)
	}
}

// Flag defines a log.Level flag with specified name, default value, and usage string.
func Flag(name string, value Level, usage string) *Level {
	l := value
	flag.CommandLine.Var(&l, name, usage)
	return &l
}
