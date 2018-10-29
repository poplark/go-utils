package log

import (
	"fmt"
)

type TerminalType int
type TerminalColor int
type ConsoleTransport struct{}

// Terminal Format
const (
	TF_DEFAULT   TerminalType = iota //  0  终端默认设置
	TF_HIGHLIGHT                     //  1  高亮显示
	TF_N0
	TF_N1
	TF_UNDERLINE //  4  使用下划线
	TF_BLINK     //  5  闪烁
	TF_N2
	TF_HIGHLIGHT2 //  7  反白显示
	TF_HIDDEN     //  8  不可见
)

// 31 - 红色前景
// 32 - 绿色前景
// 33 - 黄色前景
// 34 - 蓝色前景
// 37 - 白色前景
const (
	F_RED    TerminalColor = 31 + iota // 31
	F_GREEN                            // 32
	F_YELLOW                           // 33
	F_BLUE                             // 34
	F_N0
	F_N1
	F_WHITE // 37
)

// 40 - 黑色背景
const (
	B_BLACK TerminalColor = 40 + iota // 40
)

func NewConsoleTransport() Transport {
	var transport Transport
	consoleTransport := new(ConsoleTransport)
	transport = consoleTransport // ??? 强制转换？
	return transport
}

/*
  msg - log text: "[2000-01-01 00:00:00.000] [INFO] log text"
  logLevel - log level: [DEBUG, INFO, WARN, ERROR, FATAL]
*/
func (self *ConsoleTransport) Write(msg string, logLevel string) {
	// 如果有 log level，则根据不同的类型输出不同颜色的日志信息
	// DEBUG - 蓝色
	// INFO - 绿色
	// WARN - 黄色
	// ERROR - 红色
	// FATAL - 红色反白
	t_type := TerminalType(TF_HIGHLIGHT)
	b_color := TerminalColor(B_BLACK)
	f_color := TerminalColor(F_WHITE)
	switch logLevel {
	case "DEBUG":
		f_color = F_BLUE
		break
	case "INFO":
		f_color = F_GREEN
		break
	case "WARN":
		f_color = F_YELLOW
		break
	case "ERROR":
		f_color = F_RED
		break
	case "FATAL":
		f_color = F_RED
		t_type = TF_HIGHLIGHT2
		break
	default:
		f_color = F_WHITE
	}
	fmt.Printf("%c[%d;%d;%dm%s%c[%dm", 0x1B, t_type, b_color, f_color, msg, 0x1B, TF_DEFAULT)
}
