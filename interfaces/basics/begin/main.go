// interfaces/basics/begin/main.go
package main

import "fmt"

// define a humanoid interface with speak and walk methods returning string
type humanoid interface {
	speak()
	walk()
}

// define a person type that implements humanoid interface
type person struct {
	name string
}

func (p person) speak() {
	fmt.Printf("Person %s can speak\n", p.name)
}
func (p person) walk() {
	fmt.Printf("Person %s can walk\n", p.name)
}

// implement the Stringer interface for the person type
func (p person) Stringer() {
	fmt.Printf("Person name is %s\n", p.name)
}

// define a dog type that can walk but not speak
type dog struct {
	name string
}

func (d dog) speak() {
	fmt.Printf("Dog %s can't speak\n", d.name)
}

func (d dog) walk() {
	fmt.Printf("Dog %s can walk\n", d.name)
}

func doHumanThings(h humanoid) {
	h.speak()
	h.walk()
}

func main() {
	// invoke with a person
	doHumanThings(person{name: "Ovin"})

	// can we invoke with a dog?
	doHumanThings(dog{name: "Tabby"})

	// fmt.Println(p)
}
