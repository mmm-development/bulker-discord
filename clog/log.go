package clog

import (
	"fmt"
	"io"
	"strings"
)

type PriorityLevel int

const (
	DEBUG    PriorityLevel = 10
	INFO     PriorityLevel = 20
	WARNING  PriorityLevel = 30
	ERROR    PriorityLevel = 40
	CRITICAL PriorityLevel = 50
)

var (
	L Logger = Logger{logOuts: make([]priorityWriter, 0)}
)

func (pl PriorityLevel) String() string {
	switch pl {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case CRITICAL:
		return "CRITICAL"
	}
	return "???"
}

type priorityWriter struct {
	outstr   io.Writer
	priority PriorityLevel
}

type Logger struct {
	logOuts []priorityWriter
}

func (l *Logger) Register(stream io.Writer, priority PriorityLevel) {
	l.logOuts = append(l.logOuts, priorityWriter{
		outstr:   stream,
		priority: priority,
	})
}

func (l Logger) Log(priority PriorityLevel, msgf string, a ...interface{}) {
	msg := fmt.Sprintf("[%s] %s\n", priority, fmt.Sprintf(strings.ReplaceAll(msgf, "\n", fmt.Sprintf("\n[%s]", priority)), a...))
	for _, writer := range l.logOuts {
		if writer.priority > priority {
			continue
		}
		io.WriteString(writer.outstr, msg)
	}
}

func (l Logger) Debug(msg string, a ...interface{}) {
	l.Log(DEBUG, msg, a...)
}

func (l Logger) Info(msg string, a ...interface{}) {
	l.Log(INFO, msg, a...)
}

func (l Logger) Warning(msg string, a ...interface{}) {
	l.Log(WARNING, msg, a...)
}

func (l Logger) Error(msg string, a ...interface{}) {
	l.Log(ERROR, msg, a...)
}

func (l Logger) Critical(msg string, a ...interface{}) {
	l.Log(CRITICAL, msg, a...)
}

func (l Logger) Fatal(msg string, a ...interface{}) {
	l.Log(CRITICAL, msg, a...)
	panic(msg)
}
