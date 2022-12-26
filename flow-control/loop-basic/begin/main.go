// flow-control/loop-basic/begin/main.go
package main

import "fmt"

func main() {
	// declare a string to iterate over
	name := "Isara Ovin"

	// iterate over the string with basic for loop
	for i := 0; i < len(name); i++ {

		fmt.Println(string(name[i]))

	}
}
