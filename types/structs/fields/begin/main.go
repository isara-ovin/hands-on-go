// types/structs/fields/begin/main.go
package main

import "fmt"

// define a struct type for author
type author struct {
	firstName string
	lastName  string
}

func main() {
	// intialize author
	a := author{
		firstName: "Isara",
		lastName:  "Ovin",
	}

	// print the author
	// fmt.Printf("%#v\n", a)
	fmt.Printf("%#v", a)
}
