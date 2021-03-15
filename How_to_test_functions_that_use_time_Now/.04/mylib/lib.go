package mylib

import (
	"time"
)

type MyLib interface {
	Now() time.Time
}

type myLib struct{}

func NewMyLib() MyLib {
	return &myLib{}
}

func (m myLib) Now() time.Time {
	return time.Now()
}
