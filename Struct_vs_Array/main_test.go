package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findLoop(array []string, key string) bool {
	for _, item := range array {
		if item == key {
			return true
		}
	}

	return false
}

func createStruct(array []string) map[string]struct{} {
	result := map[string]struct{}{}
	exists := struct{}{}
	for _, item := range array {
		result[item] = exists
	}

	return result
}

func findStruct(array []string, key string) bool {
	result := createStruct(array)

	_, has := result[key]

	return has
}

func createStructMake(array []string) map[string]struct{} {
	result := make(map[string]struct{}, len(array))
	exists := struct{}{}
	for _, item := range array {
		result[item] = exists
	}

	return result
}

func findStructMake(array []string, key string) bool {
	result := createStructMake(array)

	_, has := result[key]

	return has
}

var array0 = []string{
	"keyA0", "keyA1", "keyA2", "keyA3", "keyA4", "keyA5", "keyA6", "keyA7", "keyA8", "keyA9",
	"keyB0", "keyB1", "keyB2", "keyB3", "keyB4", "keyB5", "keyB6", "keyB7", "keyB8", "keyB9",
	"keyC0", "keyC1", "keyC2", "keyC3", "keyC4", "keyC5", "keyC6", "keyC7", "keyC8", "keyC9",
	"keyD0", "keyD1", "keyD2", "keyD3", "keyD4", "keyD5", "keyD6", "keyD7", "keyD8", "keyD9",
	"keyE0", "keyE1", "keyE2", "keyE3", "keyE4", "keyE5", "keyE6", "keyE7", "keyE8", "keyE9",
	"keyF0", "keyF1", "keyF2", "keyF3", "keyF4", "keyF5", "keyF6", "keyF7", "keyF8", "keyF9",
	"keyG0", "keyG1", "keyG2", "keyG3", "keyG4", "keyG5", "keyG6", "keyG7", "keyG8", "keyG9",
	"keyH0", "keyH1", "keyH2", "keyH3", "keyH4", "keyH5", "keyH6", "keyH7", "keyH8", "keyH9",
	"keyI0", "keyI1", "keyI2", "keyI3", "keyI4", "keyI5", "keyI6", "keyI7", "keyI8", "keyI9",
	"keyJ0", "keyJ1", "keyJ2", "keyJ3", "keyJ4", "keyJ5", "keyJ6", "keyJ7", "keyJ8", "keyJ9",
	"keyK0", "keyK1", "keyK2", "keyK3", "keyK4", "keyK5", "keyK6", "keyK7", "keyK8", "keyK9",
	"keyL0", "keyL1", "keyL2", "keyL3", "keyL4", "keyL5", "keyL6", "keyL7", "keyL8", "keyL9",
	"keyM0", "keyM1", "keyM2", "keyM3", "keyM4", "keyM5", "keyM6", "keyM7", "keyM8", "keyM9",
	"keyN0", "keyN1", "keyN2", "keyN3", "keyN4", "keyN5", "keyN6", "keyN7", "keyN8", "keyN9",
	"keyO0", "keyO1", "keyO2", "keyO3", "keyO4", "keyO5", "keyO6", "keyO7", "keyO8", "keyO9",
	"keyP0", "keyP1", "keyP2", "keyP3", "keyP4", "keyP5", "keyP6", "keyP7", "keyP8", "keyP9",
}
var key = "keyP9"
var keyInexistent = "xyzfZ0"

func Benchmark_find(b *testing.B) {

	b.Run(" ForLoop          ", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = findLoop(array0, key)
		}
	})

	b.Run(" Struct           ", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = findStruct(array0, key)
		}
	})

	b.Run(" StructMake       ", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = findStructMake(array0, key)
		}
	})

	b.Run(" CreateStruct     ", func(b *testing.B) {
		result := createStruct(array0)
		has := false

		for i := 0; i < b.N; i++ {
			_, has = result[key]
		}

		_ = has
	})

	b.Run(" CreateStructMake ", func(b *testing.B) {
		result := createStructMake(array0)
		has := false

		for i := 0; i < b.N; i++ {
			_, has = result[key]
		}

		_ = has
	})
}

func Benchmark_findAssert(b *testing.B) {

	b.Run(" ForLoop          ", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			assert.True(b, findLoop(array0, key))
		}
	})

	b.Run(" Struct           ", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			assert.True(b, findStruct(array0, key))
		}
	})

	b.Run(" StructMake       ", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			assert.True(b, findStructMake(array0, key))
		}
	})

	b.Run(" CreateStruct     ", func(b *testing.B) {
		result := createStruct(array0)
		has := false

		for i := 0; i < b.N; i++ {
			_, has = result[key]
			assert.True(b, has)
		}

		_ = has
	})

	b.Run(" CreateStructMake ", func(b *testing.B) {
		result := createStructMake(array0)
		has := false

		for i := 0; i < b.N; i++ {
			_, has = result[key]
			assert.True(b, has)
		}

		_ = has
	})
}

func Benchmark_findInexistent(b *testing.B) {

	b.Run(" ForLoop          ", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			assert.False(b, findLoop(array0, keyInexistent))
		}
	})

	b.Run(" Struct           ", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			assert.False(b, findStruct(array0, keyInexistent))
		}
	})

	b.Run(" StructMake       ", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			assert.False(b, findStructMake(array0, keyInexistent))
		}
	})

	b.Run(" CreateStruct     ", func(b *testing.B) {
		result := createStruct(array0)
		has := true

		for i := 0; i < b.N; i++ {
			_, has = result[keyInexistent]
			assert.False(b, has)
		}

		_ = has
	})

	b.Run(" CreateStructMake ", func(b *testing.B) {
		result := createStructMake(array0)
		has := true

		for i := 0; i < b.N; i++ {
			_, has = result[keyInexistent]
			assert.False(b, has)
		}

		_ = has
	})
}
