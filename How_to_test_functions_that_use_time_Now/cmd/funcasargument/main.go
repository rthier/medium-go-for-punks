package main

import (
	"fmt"
	"time"
)

func whatTimeIsIt(funcTime func() time.Time) string {
	timeIs := funcTime()
	return timeIs.String()
}

func main() {

	timeNow := whatTimeIsIt(time.Now)

	fmt.Print(timeNow)

}
