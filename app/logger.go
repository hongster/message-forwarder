// Application logger
package app

import (
	"github.com/hongster/message-forwarder/logger"
	"fmt"
	"os"
)

var Logger *logger.Logger

// If `log_file` is not defined in config file, stdout is assumed.
func init() {
	logFile := Config.StringDefault("DEFAULT", "log_file", "")

	// `log_file` not defined, so assume stdout
	if logFile == "" {
		Logger = logger.NewLogger(os.Stdout)
		return
	}

	file, err := os.OpenFile(logFile, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("Unable to open log file: %v", err))
	}

	Logger = logger.NewLogger(file)
}
