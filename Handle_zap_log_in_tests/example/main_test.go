package main

import (
	"fmt"
	"sort"
	"testing"
)

func ExampleInts() {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Ints(s)
	fmt.Println(s)
	// Output: [1 2 3 4 5 6]
}

func TestInts(t *testing.T) {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Ints(s)
	fmt.Println(s)
}
