// types/slices/begin/main.go
package main

import "fmt"

func main() {
	// declare a slice string the make function
	var a = make([]string, 0)

	// append 3 names to the slice
	a = append(a, "John")
	a = append(a, "Jane")
	a = append(a, "Mary")

	// print the slice
	fmt.Println(a)
	// slice the slice using syntax slice[low:high]
	fmt.Println(a[1:3]) // [Jane Mary]
	fmt.Println(a[1:])  // [Jane Mary]
	fmt.Println(a[:1])  // [John]
	fmt.Println(a[:3])  // [John Jane Mary]
}
