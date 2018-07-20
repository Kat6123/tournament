package log

import (
	"flag"
	"fmt"
	"strconv"
)

const (
	Trace Level = iota
	Debug
	Info
	Warn
	Error
)

type Level int

func (l Level) String() string {
	switch l {
	case Trace:
		return "trace"
	case Debug:
		return "debug"
	case Info:
		return "info"
	case Warn:
		return "warn"
	case Error:
		return "error"
	default:
		return "unknown log level"
	}
}

// Set parses input string of log level and set value.
func (l *Level) Set(s string) (err error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("parse level %q as int has failed: %v", s, err)
	}

	switch Level(i) {
	case Trace:
		*l = Trace
		return nil
	case Debug:
		*l = Debug
		return nil
	case Info:
		*l = Info
		return nil
	case Warn:
		*l = Warn
		return nil
	case Error:
		*l = Error
		return nil
	default:
		return fmt.Errorf("wrong level type %q", i)
	}
	return
}

// Flag defines a log.Level flag with specified name, default value, and usage string.
func Flag(name string, value Level, usage string) *Level {
	l := value
	flag.CommandLine.Var(&l, name, usage)
	return &l
}
