// Global logger
package logger

import (
	"log"
	"os"
)

var infoLogger *log.Logger
var debugLogger *log.Logger
var warningLogger *log.Logger
var errorLogger *log.Logger

// Setup 4 tpes of logger.
func init() {
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(os.Stdout, "WARNGNG: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "ERROR", log.Ldate|log.Ltime|log.Lshortfile)
}

func INFO(format string, v ...interface{}) {
	infoLogger.Printf(format, v...)
}

func DEBUG(format string, v ...interface{}) {
	debugLogger.Printf(format, v...)
}

func WARN(format string, v ...interface{}) {
	warningLogger.Printf(format, v...)
}

func ERROR(format string, v ...interface{}) {
	errorLogger.Printf(format, v...)
}
