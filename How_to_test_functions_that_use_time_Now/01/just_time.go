package main

import (
	"fmt"
	"time"
)

func whatTimeIsItNow() string {
	timeIs := time.Now()
	return timeIs.String()
}

func main() {
	timeNow := whatTimeIsItNow()
	fmt.Print(timeNow)
}
