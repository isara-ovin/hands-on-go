// interfaces/type-switch/begin/main.go
package main

import "fmt"

// define whatAmI which takes in an argument of any type and returns inforamtion about the underlying value's type
func whatAmI(i interface{}) string {
	switch i.(type) {
	case string:
		return fmt.Sprintf("This is type %T and says %s", i, i)
	case int:
		return fmt.Sprintf("This is type %T and says %d", i, i)
	case bool:
		return fmt.Sprintf("This is type %T and says %b", i, i)
	default:
		return fmt.Sprintf("This is type %T and says %v", i, i)
	}
}
func main() {
	// invoke whatAmI function
	fmt.Println(whatAmI(1))
	fmt.Println(whatAmI("hello"))
	fmt.Println(whatAmI(true))
}
