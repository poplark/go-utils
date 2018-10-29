package main

import (
	"fmt"
	"go-utils/common/log"
)

func main() {
	fmt.Println("test logger")

	logger1 := log.NewLogger(log.INFO, ".")
	logger1.Debug("111111")
	logger1.Info("222222")
	logger1.Warn("333333")
	logger1.Error("444444")
	logger1.Fatal("555555")
}
