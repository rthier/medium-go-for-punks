package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	defer fmt.Println("end main")
	fmt.Println("start main")

	go func() {
		defer fmt.Println("end goroutine")
		fmt.Println("start goroutine")
		os.Exit(1)
	}()

	time.Sleep(10 * time.Second)

	signalChannel := make(chan os.Signal, 1)
	// signal.Ignore(os.Interrupt)
	signal.Notify(signalChannel)
	fmt.Println(signalChannel)
}
