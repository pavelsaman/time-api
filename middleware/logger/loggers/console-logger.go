package console_logger

import (
	"fmt"
	"os"

	logger_utils "github.com/pavelsaman/time-api/middleware/logger/loggers/utils"
)

type ConsoleLogger struct{}

const (
	red    = "\033[1;31m"
	orange = "\033[1;33m"
	green  = "\033[1;32m"
	end    = "\033[0m"
)

func (cl ConsoleLogger) Info(format string, v ...any) {
	fmt.Fprintf(os.Stdout, green+"INFO"+end+" ("+logger_utils.FormatUTCTime()+"): "+format, v...)
}

func (cl ConsoleLogger) Warn(format string, v ...any) {
	fmt.Fprintf(os.Stdout, orange+"WARN"+end+" ("+logger_utils.FormatUTCTime()+"): "+format, v...)
}

func (cl ConsoleLogger) Error(format string, v ...any) {
	fmt.Fprintf(os.Stderr, red+"ERROR"+end+" ("+logger_utils.FormatUTCTime()+"): "+format, v...)
}
