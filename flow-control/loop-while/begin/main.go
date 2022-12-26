// flow-control/loop-while/begin/main.go
package main

import "fmt"

func main() {
	// initialize a variable to check if the loop should continue
	x := 10

	// iterate while the condition is true
	for x > 0 {
		fmt.Printf("Now in %d\n", x)
		x--
	}
}
