package main

import (
	"go.uber.org/zap"
)

func main() {

	logger := newLogger()
	myFunction(logger)
}

func newLogger() *zap.SugaredLogger {
	return zap.NewExample().Sugar()
}

func myFunction(logger *zap.SugaredLogger) {
	logger.Info("log myFunction")
	logger.With(
		"keyOne", "valueOne",
		"keyTwo", "valueTwo",
	).Info("log with fields")
}
