// functions/begin/main.go
package main

import (
	"errors"
	"fmt"
)

// simple greet function
func greet() string {
	return "Greetings"
}

// greetWithName returns a greeting with the name
func greetWithName(name string) string {
	return fmt.Sprintf("Greetings %s", name)
}

// greetWithName returns a greeting with the name and age of the person
func greetWithNameAge(name string, age int) string {
	return fmt.Sprintf("Greetings %s, you are now %d years old", name, age)
}

// divide divides two numbers and returns the result
// if the second number is zero, it returns  error
func divide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("Canot devide by 0")
	}

	return a / b, nil
}

func main() {
	// invoke greet function
	fmt.Println(greet())

	fmt.Println(greetWithName("Toni"))

	// invoke greetWithName function
	fmt.Println(greetWithNameAge("Toni", 28))

	// invoke divide function
	fmt.Println(divide(10, 2))

	// invoke divide function with zero denominator to get an error
	fmt.Println(divide(5, 0))
}
