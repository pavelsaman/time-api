package logger_interface

type Logger interface {
	Info(message string, v ...any)
	Warn(message string, v ...any)
	Error(message string, v ...any)
}
