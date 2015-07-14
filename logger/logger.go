// Global logger
package logger

import (
	"io"
	"log"
)

type Logger struct {
	infoLogger *log.Logger
	debugLogger *log.Logger
	warningLogger *log.Logger
	errorLogger *log.Logger
}

func NewLogger(writer io.Writer) (logger *Logger) {
	return &Logger{
		infoLogger: log.New(writer, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		debugLogger: log.New(writer, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		warningLogger: log.New(writer, "WARNGNG: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(writer, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (logger *Logger) Info(format string, v ...interface{}) {
	logger.infoLogger.Printf(format, v...)
}

func (logger *Logger) Debug(format string, v ...interface{}) {
	logger.debugLogger.Printf(format, v...)
}

func (logger *Logger) Warn(format string, v ...interface{}) {
	logger.warningLogger.Printf(format, v...)
}

func (logger *Logger) Error(format string, v ...interface{}) {
	logger.errorLogger.Printf(format, v...)
}
