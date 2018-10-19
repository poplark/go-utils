package log

import (
	"fmt"
	"os"
	"time"
)

type LoggerLevel int

type Logger struct {
	level   LoggerLevel
	output  string
	logfile *os.File
	// buf - text buffer size
}

const (
	DEBUG LoggerLevel = 1 + iota
	INFO
	WARN
	ERROR
	// panic
	FATAL
)

func NewLogger(level LoggerLevel, output string) *Logger {
	fmt.Println("config", level, output)
	logger := Logger{level, output, nil}
	logger.OpenLogFile()
	return &logger
}

func (logger *Logger) Debug(msg string) {
	if logger.level > DEBUG {
		return
	}
	logger.Write("DEBUG", msg)
}
func (logger *Logger) Info(msg string) {
	if logger.level > INFO {
		return
	}
	logger.Write("INFO", msg)
}
func (logger *Logger) Warn(msg string) {
	if logger.level > WARN {
		return
	}
	logger.Write("WARN", msg)
}
func (logger *Logger) Error(msg string) {
	if logger.level > ERROR {
		return
	}
	logger.Write("ERROR", msg)
}
func (logger *Logger) Fatal(msg string) {
	if logger.level > FATAL {
		return
	}
	logger.Write("FATAL", msg)
}

func (logger *Logger) OpenLogFile() error {
	filepath := fmt.Sprintf("%s/%s", logger.output, "test.log")
	logfile, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	logger.logfile = logfile
	return nil
}
func (logger *Logger) Write(t string, msg string) {
	// todo - use fmt as a transport
	now := time.Now()
	txt := fmt.Sprintf("[%s] [%s] - %s\n", now.Format("2006-01-02 15:04:05.000"), t, msg)
	log := []byte(txt)
	// todo - deal with error
	logger.logfile.Write(log)
}
