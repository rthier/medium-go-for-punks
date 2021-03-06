package main

import (
	"go.uber.org/zap"
)

func main() {

	logger := zap.NewExample().Sugar()
	myFunction(logger)
}

func myFunction(logger *zap.SugaredLogger) {
	logger.Info("log myFunction")
	logger.With(
		"keyOne", "valueOne",
		"keyTwo", "valueTwo",
	).Info("log with fields")
}
