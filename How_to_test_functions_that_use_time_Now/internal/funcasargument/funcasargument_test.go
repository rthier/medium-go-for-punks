package funcasargument

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_whatTimeIsIt(t *testing.T) {

	t.Run("Return a string with now time", func(t *testing.T) {
		// Given
		mockNow := func() time.Time {
			fakeTime, _ := time.Parse(time.RFC3339, "2021-01-08T15:04:05Z")
			return fakeTime
		}

		// When
		got := WhatTimeIsIt(mockNow)

		// Then
		assert.Equal(t, "2021-01-08 15:04:05 +0000 UTC", got)
	})
}
