package main

import (
	"fmt"
	"time"
)

type funcTimeType func() time.Time

func whatTimeIsIt(funcTime funcTimeType) string {
	return funcTime().String()
}

func main() {
	fmt.Print(whatTimeIsIt(time.Now))
}
