package log

import (
	"fmt"
	"os"
	"time"
)

type FileTransport struct {
	output     string
	prefix     string
	suffix     string
	logFiles   map[string]*os.File
	buf        chan []byte // for accumulating text to write
	isClosed   chan bool   // for accumulating text to write
	isDetached bool
}

const BUFFER_LEN = 1000

func NewFileTransport(output, prefix, suffix string, isDetached bool, size int64) Transport {
	if output == "" {
		output = "."
	}

	if output[len(output)-1:] == "/" {
		output = output[:len(output)-1]
	}
	var transport Transport
	fileTransport := new(FileTransport) // {output, nil, nil, nil}

	fileTransport.output = output
	fileTransport.prefix = prefix
	fileTransport.suffix = suffix
	fileTransport.logFiles = make(map[string]*os.File)
	if isDetached == true {
		outFilepath := LogName(output, prefix, "out", suffix, 0)
		errorFilepath := LogName(output, prefix, "error", suffix, 0)
		outFile, _ := os.OpenFile(outFilepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
		errorFile, _ := os.OpenFile(errorFilepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
		fileTransport.logFiles["out"] = outFile
		fileTransport.logFiles["error"] = errorFile
	} else {
		defaultFilepath := LogName(output, prefix, "", suffix, 0)
		defaultFile, _ := os.OpenFile(defaultFilepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
		fileTransport.logFiles["default"] = defaultFile
	}

	if size > 0 {
		fileTransport.buf = make(chan []byte, size)
	} else {
		fileTransport.buf = make(chan []byte, BUFFER_LEN)
	}
	fileTransport.isDetached = isDetached
	// fileTransport.logfile = nil
	// fileTransport.buf = nil
	// fileTransport.isClosed = nil
	transport = fileTransport // 强制转换？
	return transport
}

func LogName(dir, prefix, tp, suffix string, idx int64) string {
	t := time.Now()
	return fmt.Sprintf("%s/%s%d%02d%02d-%s%d%s", dir, prefix, t.Year(), t.Month(), t.Day(), tp, idx, suffix)
}

/*
  msg - log text: "[2000-01-01 00:00:00.000] [INFO] log text"
  logLevel - log level: [DEBUG, INFO, WARN, ERROR, FATAL]
*/
func (self *FileTransport) Write(msg string, logLevel string) {
	// 如果有 log level，则根据不同的类型输出到不同日志文件中
	// 如果没有 log level，则将 log text 写至统一的日志文件中
	// DEBUG, INFO - xx_20181010-out0.log
	// WARN, ERROR, FATAL - xx_20181010-error0.log
	log := []byte(msg)
	if self.isDetached {
		switch logLevel {
		case "DEBUG":
			self.logFiles["out"].Write(log) // 不可省去 break ?
			break
		case "INFO":
			self.logFiles["out"].Write(log)
			break
		case "WARN":
			self.logFiles["error"].Write(log)
			break
		case "ERROR":
			self.logFiles["error"].Write(log)
			break
		case "FATAL":
			self.logFiles["error"].Write(log)
			break
		default:
			self.logFiles["out"].Write(log)
		}
	} else {
		self.logFiles["default"].Write(log)
	}
}
