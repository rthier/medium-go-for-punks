package main

import (
	"go.uber.org/zap"
)

func main() {

	logger := zap.NewExample()
	myFunction(logger)
}

func myFunction(logger *zap.Logger) {
	logger.Info("log myFunction")
}
