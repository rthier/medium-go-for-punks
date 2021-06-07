package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func Test_zap(t *testing.T) {

	t.Run("Handle log message", func(t *testing.T) {
		// Given
		observedZapCore, observedLogs := observer.New(zap.InfoLevel)
		observedLogger := zap.New(observedZapCore)

		// When
		myFunction(observedLogger)

		// Then
		require.Equal(t, 1, observedLogs.Len())
		firstLog := observedLogs.All()[0]
		assert.Equal(t, "log myFunction", firstLog.Message)

	})
}
