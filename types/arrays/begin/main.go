// types/arrays/begin/main.go
package main

import "fmt"

func main() {
	// declare an array of integers
	// var a [3]int
	// a := [3]int{0, 0, 0} // or without it, defaulrs to 0s
	var a = [3]int{1, 2, 3}
	// print the array
	fmt.Println(a)
	// set the first element to 1
	a[1] = 1
	// print the array
	fmt.Println(a)

}
