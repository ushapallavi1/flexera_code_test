package logger

import "log"

type Logger interface {
	Info(v ...any)
	Error(v ...any)
}

var _ = new(logger)

type logger struct{}

func NewLogger() Logger {
	return &logger{}
}

func (l *logger) Info(v ...any) {
	_log := log.Default()
	_log.SetPrefix("[INFO] ")
	_log.Println(v...)
}

func (l *logger) Error(v ...any) {
	_log := log.Default()
	_log.SetPrefix("[Error] ")
	_log.Println(v...)
}
