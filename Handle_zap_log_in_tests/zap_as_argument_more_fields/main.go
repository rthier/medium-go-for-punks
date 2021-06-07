package main

import (
	"go.uber.org/zap"
)

func main() {

	logger := newLogger()
	myFunction(logger)
}

func newLogger() *zap.Logger {
	return zap.NewExample()
}

func myFunction(logger *zap.Logger) {
	logger.Info("log myFunction")
	logger.With(
		zap.Field{Key: "keyOne", String: "valueOne"},
		zap.Field{Key: "keyTwo", String: "valueTwo"},
	).Info("log with fields")
}
