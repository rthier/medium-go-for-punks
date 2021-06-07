package main

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

func Test_zap(t *testing.T) {

	t.Run("Handle log message with fields", func(t *testing.T) {
		// Given
		// observedZapCore, observedLogs := observer.New(zap.InfoLevel)
		// observedLoggerSugared := zap.New(observedZapCore).Sugar()

		// When
		myFunction(zaptest.NewLogger(t).Sugar())
		myFunction(zap.NewExample().Sugar())

		// Then
		// require.Equal(t, 2, observedLogs.Len())
		// allLogs := observedLogs.All()
		// assert.Equal(t, "log myFunction", allLogs[0].Message)
		// assert.Equal(t, "log with fields", allLogs[1].Message)
		// assert.ElementsMatch(t, []zap.Field{
		// 	zap.String("keyOne", "valueOne"),
		// 	zap.String("keyTwo", "valueTwo"),
		// }, allLogs[1].Context)
	})
}
