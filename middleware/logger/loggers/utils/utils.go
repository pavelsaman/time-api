package logger_utils

import "time"

func FormatUTCTime() string {
	return time.Now().UTC().Format(time.RFC3339Nano)
}
