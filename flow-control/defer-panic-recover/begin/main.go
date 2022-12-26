// flow-control/defer-panic-recover/begin/main.go
package main

import "fmt"

func cleanup(msg string) {
	fmt.Println(msg)
}

func main() {

	fmt.Println("Testing in main")
	// defer function calls
	defer cleanup("first")
	defer cleanup("second")
	// defer recovery
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from panic:", err)
		}
	}()

	// panic
	panic("Dummy panic")
}
