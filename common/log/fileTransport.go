package log

import (
	"os"
)

type FileTransport struct {
	output   string
	logfile  *os.File
	buf      chan []byte // for accumulating text to write
	isClosed chan bool   // for accumulating text to write
}

func NewFileTransport(output, prefix, suffix string, size int64) Transport {
	if output == "" {
		output = "."
	}

	if output[len(output)-1:] == "/" {
		output = output[:len(output)-1]
	}
	var transport Transport
	fileTransport := new(FileTransport) // {output, nil, nil, nil}
	// fileTransport.output = output
	// fileTransport.logfile = nil
	// fileTransport.buf = nil
	// fileTransport.isClosed = nil
	transport = fileTransport // 强制转换？
	return transport
}

/*
  t - log type: [DEBUG, INFO, WARN, ERROR, FATAL]
  msg - log text: "[2000-01-01 00:00:00.000] [INFO] log text"
*/
func (self *FileTransport) Write(t string, msg string) {
	// 如果有 log type，则将 log text 分别写至不同类型的日志文件中
	// 如果没有 log type，则将 log text 写至统一的日志文件中
	log := []byte(msg)
	self.logfile.Write(log)
}
