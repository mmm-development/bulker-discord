package log

import (
	"fmt"
	"io"
)

type PriorityLevel int

const (
	DEBUG    PriorityLevel = 10
	INFO     PriorityLevel = 20
	WARNING  PriorityLevel = 30
	ERROR    PriorityLevel = 40
	CRITICAL PriorityLevel = 50
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

func NewLogger() Logger {
	return Logger{
		logOuts: make([]priorityWriter, 0),
	}
}

func (l *Logger) Register(stream io.Writer, priority PriorityLevel) {
	l.logOuts = append(l.logOuts, priorityWriter{
		outstr:   stream,
		priority: priority,
	})
}

func (l Logger) Log(priority PriorityLevel, msg string) {
	for _, writer := range l.logOuts {
		if writer.priority > priority {
			continue
		}
		io.WriteString(writer.outstr, fmt.Sprintf("[%s] %s", priority, msg))
	}
}

func (l Logger) Info(msg string) {
	l.Log(INFO, msg)
}

func (l Logger) Warning(msg string) {
	l.Log(WARNING, msg)
}

func (l Logger) Error(msg string) {
	l.Log(ERROR, msg)
}

func (l Logger) Critical(msg string) {
	l.Log(CRITICAL, msg)
}

func (l Logger) Fatal(msg string) {
	l.Log(CRITICAL, msg)
	panic(msg)
}
