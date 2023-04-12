package console_logger

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestLogFunctions(t *testing.T) {
	consoleLogger := ConsoleLogger{}

	funcs := map[string]func(io.Writer, string, ...any){
		"INFO":  consoleLogger.Info,
		"WARN":  consoleLogger.Warn,
		"ERROR": consoleLogger.Error,
	}

	for ft, f := range funcs {
		buffer := &bytes.Buffer{}
		f(buffer, "%s", "test")

		if !strings.Contains(buffer.String(), ft) {
			t.Errorf("Expected \"%s\" in \"%s\"", ft, buffer.String())
		}

		if !strings.Contains(buffer.String(), "test") {
			t.Errorf("Expected \"test\" in \"%s\"", buffer.String())
		}
	}
}
