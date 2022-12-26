// flow-control/loop-range/begin/main.go
package main

import "fmt"

func main() {
	// initialize a slice of ints
	x := []int{1, 2, 3, 4, 5}

	// use for-range to iterate over the slice and print each value
	for i, v := range x {
		fmt.Printf("Currently in index %d, value %d\n", i, v)
	}

	// declare a map of strings to ints
	y := map[string]int{"A": 1, "B": 2, "C": 3}

	// use for-range to iterate over the map and print each key/value pair
	for k, v := range y {
		fmt.Printf("Currently in key %s, value %d\n", k, v)
	}
}
