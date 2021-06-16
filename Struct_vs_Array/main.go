package main

import (
	"fmt"
)

// func findLoop(array []string, key string) bool {
// 	for _, item := range array {
// 		if item == key {
// 			return true
// 		}
// 	}

// 	return false
// }

// func findStruct(array []string, key string) bool {
// 	result := map[string]struct{}{}
// 	exists := struct{}{}
// 	for _, item := range array {
// 		result[item] = exists
// 	}

// 	_, has := result[key]

// 	return has
// }

func main() {

	array0 := []string{"KEY1", "KEY2", "KEY3"}

	fmt.Println(array0)

	// fmt.Println(findLoop(array0, "KEY2"))
	// fmt.Println(findStruct(array0, "KEY2"))

}
