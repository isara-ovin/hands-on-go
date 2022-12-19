// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "global"

func printVal() {
	fmt.Println(val)
}

func updateGlobal() {
	val = "global updated"
	fmt.Println(val)
}

func main() {
	fmt.Println(val)
	// create a local variable "val" and assign it an int value
	val := 10

	// print the value of the local variable "val"
	fmt.Println(val)

	// print the value of the package-level variable "val"
	func() {
		fmt.Println("On Anonymous function", val)
	}()

	printVal()

	// update the package-level variable "val" and print it again
	updateGlobal()
	// print the pointer address of the local variable "val"

	fmt.Printf("%T, local &val = %v\n", &val, &val)

	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*(&val) = 30

	fmt.Println("After changing pointer value", val)
}
