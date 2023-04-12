package console_logger

import (
	"fmt"
	"os"
)

type ConsoleLogger struct{}

const (
	red    = "\033[1;31m"
	orange = "\033[1;33m"
	green  = "\033[1;32m"
	end    = "\033[0m"
)

func (cl ConsoleLogger) Info(format string, v ...any) {
	fmt.Fprintf(os.Stdout, green+"INFO: "+end+format, v...)
}

func (cl ConsoleLogger) Warn(format string, v ...any) {
	fmt.Fprintf(os.Stdout, orange+"WARN: "+end+format, v...)
}

func (cl ConsoleLogger) Error(format string, v ...any) {
	fmt.Fprintf(os.Stderr, red+"ERROR: "+end+format, v...)
}
