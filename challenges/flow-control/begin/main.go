// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// handle any panics that might occur anywhere in this function
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("Error occured")
	// 	}
	// }()
	var data *[]byte
	// use os package to read the file specified as a command line argument
	args := os.Args

	if len(args) != 2 {
		panic(fmt.Errorf("Error in reading args"))
	} else {

		path := args[1]
		file, err := os.ReadFile(path)
		check(err)

		data = &file

	}
	// convert the bytes to a string
	dataStream := string(*data)
	// initialize a map to store the counts
	counts := map[string]int{"letters": 0, "symbols": 0, "numbers": 0, "words": 0}

	// use the standard library utility package that can help us split the string into words
	words := strings.Split(dataStream, " ")

	// capture the length of the words slice
	counts["words"] = len(words)

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	for _, w := range words{
		
	}

	// dump the map to the console using the spew package
	//
}
