package main

import (
	"fmt"
	"go-utils/common/log"
)

func main() {
	fmt.Println("test logger")

	fmt.Println("test logger", log.INFO)
	logger1 := log.NewLogger(log.INFO, ".")
	logger1.Info("1111")
	logger1.Info("2222")
	logger1.Info("6666")
}
