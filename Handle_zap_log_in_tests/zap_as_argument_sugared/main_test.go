package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func Test_zap(t *testing.T) {

	t.Run("Handle log message with fields", func(t *testing.T) {
		// Given
		observedZapCore, observedLogs := observer.New(zap.InfoLevel)
		observedLoggerSugared := zap.New(observedZapCore).Sugar()

		// When
		myFunction(observedLoggerSugared)

		// Then
		require.Equal(t, 2, observedLogs.Len())
		allLogs := observedLogs.All()
		assert.Equal(t, "log myFunction", allLogs[0].Message)
		assert.Equal(t, "log with fields", allLogs[1].Message)
		assert.ElementsMatch(t, []zap.Field{
			zap.String("keyOne", "valueOne"),
			zap.String("keyTwo", "valueTwo"),
		}, allLogs[1].Context)
	})
}
