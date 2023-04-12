package logger_interface

import "io"

type Logger interface {
	Info(writer io.Writer, message string, v ...any)
	Warn(writer io.Writer, message string, v ...any)
	Error(writer io.Writer, message string, v ...any)
}
