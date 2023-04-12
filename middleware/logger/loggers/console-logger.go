package console_logger

import (
	"fmt"
	"io"

	logger_utils "github.com/pavelsaman/time-api/middleware/logger/loggers/utils"
)

type ConsoleLogger struct{}

const (
	red    = "\033[1;31m"
	orange = "\033[1;33m"
	green  = "\033[1;32m"
	end    = "\033[0m"
)

func (cl ConsoleLogger) Info(writer io.Writer, format string, v ...any) {
	fmt.Fprintf(writer, green+"INFO"+end+" ("+logger_utils.FormatUTCTime()+"): "+format, v...)
}

func (cl ConsoleLogger) Warn(writer io.Writer, format string, v ...any) {
	fmt.Fprintf(writer, orange+"WARN"+end+" ("+logger_utils.FormatUTCTime()+"): "+format, v...)
}

func (cl ConsoleLogger) Error(writer io.Writer, format string, v ...any) {
	fmt.Fprintf(writer, red+"ERROR"+end+" ("+logger_utils.FormatUTCTime()+"): "+format, v...)
}
