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
	logger.With(
		zap.Field{Key: "keyOne", String: "valueOne"},
		zap.Field{Key: "keyTwo", String: "valueTwo"},
	).Info("log with fields")
}
