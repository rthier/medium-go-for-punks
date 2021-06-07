package main

import (
	"go.uber.org/zap"
)

func main() {
	ExampleLog()
}

func ExampleLog() {

	logger := zap.NewExample()
	logger.Info("log example")
	// Output: {"level":"info","msg":"log example"}
}
