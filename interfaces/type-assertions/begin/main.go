// interfaces/type-assertions/begin/main.go
package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) sayMyName() {
	fmt.Printf("Say my name %s\n", p.name)
}

func (p person) sayMyAge() {
	fmt.Printf("My age is %d\n", p.age)
}

type anyName interface {
	sayMyName()
}

func name(a anyName) {
	ta := a.(person) // we can call recivers that are not included in the
	ta.sayMyAge()

	a.sayMyName()

	a.(person).sayMyAge() // same assertion for readability
}

func main() {
	// perform a type assertion
	p := person{name: "Ovin", age: 28}
	name(p)
	// perform a type assertion and handle the error
}
