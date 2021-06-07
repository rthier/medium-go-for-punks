package main

import (
	"fmt"

	"github.com/rthier/medium-go-for-punks/How_to_test_functions_that_use_time_Now/How_to_test_functions_that_use_time_Now/internal/structandinterface"
)

func main() {

	myLib := structandinterface.NewMyLib()
	fmt.Print(myLib.Now())
}
