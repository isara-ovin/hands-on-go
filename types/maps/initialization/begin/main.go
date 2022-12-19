// types/maps/initialization/begin/main.go
package main

import "fmt"

type author struct {
	name string
}

func main() {
	// declare a map of string keys and author values
	a := make(map[string]author)

	// initialize the map with make
	//

	// add authors to the map
	a["ovin"] = author{name: "ovin"}
	a["isara"] = author{name: "isara"}
	a["savini"] = author{name: "savini"}

	// print the map with fmt.Printf
	fmt.Printf("%#v\n", a)

	// read a value from the map with a known key
	delete(a, "isara")
	val, ok := a["isara"]

	fmt.Println(val, ok)

	fmt.Printf("%#v\n", a)

}
