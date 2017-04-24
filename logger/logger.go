package logger

import (
	"io"
	"log"
)

// Logger interface for the logger. used to abstract away the underlying log interface
type Logger interface {
	Errorf(format string, args ...interface{})
}

//Log implements Logger.
type Log struct {
	*log.Logger
}

//New Creates a new instance of Log
func New(out io.Writer) *Log {
	logger := log.New(out, "[naa]", log.LUTC)
	return &Log{logger}
}

//Errorf is the interface method
func (l *Log) Errorf(format string, args ...interface{}) {
	l.Printf(format, args)
}
